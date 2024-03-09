package auth

import (
	"awesomeProject1/dbconnect"
	"context"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func ShowWaitingList(w http.ResponseWriter, r *http.Request) {
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

	// Проверка метода запроса
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		glog.Fatal("http.StatusMethodNotAllowed")
	}

	// Подключение к базе данных
	client := dbconnect.ConnectToDB()
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			glog.Fatal(err)
		}
	}()
	coll := client.Database("orphanage").Collection("waitinglist")
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	var result []interface{}
	for cursor.Next(context.Background()) {
		var cur map[string]interface{}
		if err = cursor.Decode(&cur); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		result = append(result, cur)
	}
	if err = cursor.Close(context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	glog.Info("Showed successfully")
	// Отправка данных в формате JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
}
