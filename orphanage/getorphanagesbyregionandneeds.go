package orphanage

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetOrphanagesByRegionAndNeeds(w http.ResponseWriter, r *http.Request) {
	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	orphanageColl := client.Database("orphanage").Collection("orphanage")
	needsColl := client.Database("orphanage").Collection("need")
	// Парсинг данных из тела запроса
	var orphanageFilter structures.OrphanageFilter
	if err := json.NewDecoder(r.Body).Decode(&orphanageFilter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	needCursor, err := needsColl.Find(context.Background(), bson.M{"categoryofdonate": orphanageFilter.CategoryOfDonate})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var result []interface{}
	for needCursor.Next(context.Background()) {
		var need map[string]interface{}
		if err = needCursor.Decode(&need); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		objId, err := primitive.ObjectIDFromHex(need["orphanageid"].(string))
		if err != nil {
			// Обработка ошибки при парсинге строки в ObjectID
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//
		var orphanage map[string]interface{}
		err = orphanageColl.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&orphanage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if orphanage["region"] == orphanageFilter.Region {
			result = append(result, need)
		}
	}
	err = needCursor.Close(context.Background())
	if err != nil {
		return
	}
	if result == nil {
		needCursor, err = needsColl.Find(context.Background(), bson.M{"categoryofdonate": orphanageFilter.CategoryOfDonate})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for needCursor.Next(context.Background()) {
			var need map[string]interface{}
			if err = needCursor.Decode(&need); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			result = append(result, need)
		}
		err = needCursor.Close(context.Background())
		if err != nil {
			return
		}
	}
	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}
