package orphanage

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func EditOrphanage(w http.ResponseWriter, r *http.Request) {
	var err error

	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Парсинг данных из тела запроса
	var orphanageData structures.Orphanage
	if err := json.NewDecoder(r.Body).Decode(&orphanageData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	coll := client.Database("orphanage").Collection("orphanage")

	filter := bson.D{{"_id", orphanageData.ID}}

	update := bson.D{{"$set", orphanageData}} //изменить все поля, если нет какого-то поля то добавить

	_, err = coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}

	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Successfully edited")
	if err != nil {
		return
	}
}
