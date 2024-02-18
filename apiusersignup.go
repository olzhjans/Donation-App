package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func apiUserSignUp(w http.ResponseWriter, r *http.Request) {
	var err error

	client := connectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("orphanage").Collection("users")

	// Парсинг данных из тела запроса
	var user Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//currentTime := time.Now()

	// Вставка данных в базу данных
	_, err = coll.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Возвращаем успешный статус
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Added successfully")
	if err != nil {
		return
	}
}
