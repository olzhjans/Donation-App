package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func apiUserLogin(w http.ResponseWriter, r *http.Request) {
	var err error
	client := connectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Парсинг данных из тела запроса
	var loginData LoginData
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userColl := client.Database("orphanage").Collection("users")
	adminsColl := client.Database("orphanage").Collection("admins")
	var result bson.M
	err = userColl.FindOne(context.TODO(), bson.D{{"phone", loginData.Phone}, {"password", loginData.Password}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = adminsColl.FindOne(context.TODO(), bson.D{{"phone", loginData.Phone}, {"password", loginData.Password}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			result = bson.M{"result": "No documents"}
		}
	}
	/*if err != nil {
		panic(err)
	}*/

	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}
