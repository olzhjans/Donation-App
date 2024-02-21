package orphanage

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"net/http"
)

func AddOrphanage(w http.ResponseWriter, r *http.Request) {
	var err error
	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("orphanage").Collection("orphanage")

	// Парсинг данных из тела запроса
	var orphanage structures.Orphanage
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Successfully added")
	if err != nil {
		return
	}
}
