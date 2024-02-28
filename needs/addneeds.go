package needs

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func AddNeeds(w http.ResponseWriter, r *http.Request) {
	var err error
	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("orphanage").Collection("need")
	// Парсинг данных из тела запроса
	var need structures.Need
	if err := json.NewDecoder(r.Body).Decode(&need); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	need.Expiring = primitive.NewDateTimeFromTime(time.Now().UTC().AddDate(0, 1, 0))
	// Вставка данных в базу данных
	_, err = coll.InsertOne(context.Background(), need)
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
