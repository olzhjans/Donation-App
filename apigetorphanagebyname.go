package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func apiGetByName(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получение ID коллекции из URL
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Подключение к базе данных
	client := connectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("orphanage").Collection("orphanage")

	// Поиск данных по ID
	var orphanage Orphanage
	err := coll.FindOne(context.Background(), bson.M{"name": name}).Decode(&orphanage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка данных в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orphanage)
}
