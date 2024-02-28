package auth

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/mail"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func DeleteWaitingListById(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodDelete {
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
	// Получение ID записи из запроса
	params := r.URL.Query()
	id := params.Get("id")
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}
	// Преобразование ID в формат BSON
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	filter := bson.M{"_id": objID}
	_, err = coll.DeleteOne(context.Background(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Отправка ответа об успешном удалении
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Deleted successfully")
	if err != nil {
		return
	}
	//SEND MAIL
	mail.SendMail("olzhjans@gmail.com", "Donation-App", "Unfortunately, your request has been rejected")
}
