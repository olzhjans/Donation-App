package auth

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/mail"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func ConfirmRegistrationById(w http.ResponseWriter, r *http.Request) {
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
	waitingColl := client.Database("orphanage").Collection("waitinglist")

	// Получение ID коллекции из URL
	id := r.URL.Query().Get("_id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Поиск данных по имени
	var admin structures.Admins
	err = waitingColl.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	adminColl := client.Database("orphanage").Collection("admins")
	_, err = adminColl.InsertOne(context.Background(), admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//DELETE FROM WAITING LIST
	_, err = waitingColl.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//успешно
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Confirmed successfully")
	if err != nil {
		return
	}
	//SEND MAIL
	mail.SendMail("olzhjans@gmail.com", "Donation-App", "Congratulations, your request has been confirmed")
}
