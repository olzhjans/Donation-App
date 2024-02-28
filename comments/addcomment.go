package comments

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	var err error
	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	commentaryColl := client.Database("orphanage").Collection("comments")
	var commentary structures.Commentary
	if err := json.NewDecoder(r.Body).Decode(&commentary); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	commentary.Date = primitive.NewDateTimeFromTime(time.Now().UTC())
	_, err = commentaryColl.InsertOne(context.Background(), commentary)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Successfully added")
	if err != nil {
		return
	}
}
