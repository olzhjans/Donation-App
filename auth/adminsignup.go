package auth

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"net/http"
)

func AdminSignUp(w http.ResponseWriter, r *http.Request) {
	var err error

	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("orphanage").Collection("admins")

	// Парсинг данных из тела запроса
	var admin structures.Admins
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//currentTime := time.Now()

	// Вставка данных в базу данных
	_, err = coll.InsertOne(context.Background(), admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный статус и сообщение об успешном добавлении
	//successMessage := map[string]string{"message": "Данные успешно добавлены в базу"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Added successfully")
	if err != nil {
		return
	}
}
