package auth

import (
	"awesomeProject1/dbconnect"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func ShowWaitingList(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Подключение к базе данных
	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("orphanage").Collection("waitinglist")
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println(err)
		return
	}
	var result []interface{}
	for cursor.Next(context.Background()) {
		var cur map[string]string
		if err := cursor.Decode(&cur); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result = append(result, cur)
	}
	// Отправка данных в формате JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)

	if err := cursor.Close(context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
