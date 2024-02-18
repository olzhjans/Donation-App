package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func apiAddNewOrphanage(w http.ResponseWriter, r *http.Request) {
	var err error
	client := connectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("orphanage").Collection("orphanage")

	// Парсинг данных из тела запроса
	var orphanage Orphanage
	if err := json.NewDecoder(r.Body).Decode(&orphanage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вставка данных в базу данных
	_, err = coll.InsertOne(context.Background(), orphanage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный статус
	w.WriteHeader(http.StatusCreated)
}
