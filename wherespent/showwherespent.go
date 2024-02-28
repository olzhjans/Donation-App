package wherespent

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func ShowWhereSpent(w http.ResponseWriter, r *http.Request) {
	// Подключение к базе данных
	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("orphanage").Collection("wherespent")
	var whereSpentFilter structures.WhereSpentFilter
	if err := json.NewDecoder(r.Body).Decode(&whereSpentFilter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cursor, err := coll.Find(context.Background(), bson.M{"orphanageid": whereSpentFilter.OrphanageId, "date": bson.M{"$gte": whereSpentFilter.From, "$lte": whereSpentFilter.To}})
	if err != nil {
		panic(err)
	}
	/*
		if orphanageid == "" {
			cursor, err = coll.Find(context.Background(), bson.D{})
			if err != nil {
				panic(err)
			}
		}
	*/
	var result []interface{}
	for cursor.Next(context.Background()) {
		var cur map[string]interface{}
		if err := cursor.Decode(&cur); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result = append(result, cur)
	}
	// Отправка данных в формате JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err := cursor.Close(context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
