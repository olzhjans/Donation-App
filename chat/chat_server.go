package chat

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

// The ClientManager will keep track of all the connected clients, clients that are trying to become registered, clients that have
// become destroyed and are waiting to be removed, and messages that are to be broadcasted to and from all connected clients.
type ClientManager struct {
	clients       map[*Client]bool
	broadcast     chan []byte
	directMessage chan []byte
	register      chan *Client
	unregister    chan *Client
}

type Client struct { // Each Client has a unique id, a socket connection, and a message waiting to be sent.
	id     string
	socket *websocket.Conn
	send   chan []byte
}

type Message struct {
	Sender    string `json:"sender,omitempty"`    // отправитель
	Recipient string `json:"recipient,omitempty"` // получатель
	Content   string `json:"content,omitempty"`   // текст
}

var manager = ClientManager{
	broadcast:     make(chan []byte),
	directMessage: make(chan []byte),
	register:      make(chan *Client),
	unregister:    make(chan *Client),
	clients:       make(map[*Client]bool),
}

func (manager *ClientManager) start() {
	fmt.Println(">>start")
	var err error
	// DB CONNECT
	client := dbconnect.ConnectToDB()
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			glog.Fatal(err)
		}
	}()
	chatColl := client.Database("orphanage").Collection("chat")
	for {
		select {
		case conn := <-manager.register:
			fmt.Println(">>>manager.register")
			fmt.Println("user connected ", conn)
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			fmt.Println(">>>manager.unregister")
			fmt.Println("user disconnected ", conn)
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			fmt.Println(">>>manager.broadcast")
			for conn := range manager.clients {
				fmt.Println(conn)
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		case jsonMessage := <-manager.directMessage:
			fmt.Println(">>>manager.directMessage")
			// Получаем JSON-представление сообщения из канала
			var msg Message
			if err = json.Unmarshal(jsonMessage, &msg); err != nil {
				fmt.Println("Error unmarshalling message:", err)
				continue
			}
			// Отправляем сообщение конкретному получателю
			recipientID := msg.Recipient
			for conn := range manager.clients {
				if conn.id == recipientID {
					conn.send <- jsonMessage // Отправляем JSON-представление сообщения получателю
				}
			}
			// SAVE MESSAGE TO DB
			// Prepare document to insert
			doc := structures.Chat{
				Sender:    msg.Sender,
				Recipient: msg.Recipient,
				Content:   msg.Content,
				Date:      primitive.NewDateTimeFromTime(time.Now().Add(5 * time.Hour)),
			}
			// Вставка данных в базу данных
			_, err = chatColl.InsertOne(context.Background(), doc)
			if err != nil {
				glog.Fatal(err)
			}
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	fmt.Println(">>send")
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}

func (c *Client) read() { // The point of this goroutine is to read the socket data and add it to the manager.broadcast for further orchestration.
	fmt.Println(">>read")
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()
	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			//manager.unregister <- c
			c.socket.Close()
			break
		}
		// ADD SENDER
		var updatedMessage map[string]interface{}
		// Десериализуем JSON
		if err = json.Unmarshal(message, &updatedMessage); err != nil {
			glog.Fatal(err)
		}
		result, _ := json.Marshal(&Message{Sender: c.id, Recipient: updatedMessage["recipient"].(string), Content: updatedMessage["content"].(string)})
		manager.directMessage <- result // Отправляем сообщение в канал
	}
}

func (c *Client) write() {
	fmt.Println(">>write")
	defer func() {
		c.socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func wsPage(res http.ResponseWriter, req *http.Request) {
	fmt.Println(">>wsPage")
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	client := &Client{id: uuid.NewV4().String(), socket: conn, send: make(chan []byte)}
	manager.register <- client
	go client.read()
	go client.write()
}

func LaunchChatServer() {
	fmt.Println(">>LaunchChatServer")
	var err error
	err = flag.Set("logtostderr", "false") // Логировать в stderr (консоль) (false для записи в файл)
	if err != nil {
		log.Fatal(err)
	}
	err = flag.Set("stderrthreshold", "FATAL") // Устанавливаем порог для вывода ошибок в stderr
	if err != nil {
		log.Fatal(err)
	}
	err = flag.Set("log_dir", "C:/golang/logs/") // Указываем директорию для сохранения логов
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()
	defer glog.Flush()
	fmt.Println("Starting chat server...")
	go manager.start()
	http.HandleFunc("/ws", wsPage)
	go func() {
		glog.Fatal(http.ListenAndServe(":12345", nil))
	}()
}
