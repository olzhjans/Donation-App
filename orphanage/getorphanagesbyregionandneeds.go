package orphanage

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"fmt"
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
	orphanageCursor, err := orphanageColl.Find(context.Background(), bson.M{"region": orphanageFilter.Region})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//IF FOUND NOTHING THEN FIND ALL REGIONS
	if orphanageCursor.RemainingBatchLength() == 0 {
		orphanageCursor, err = orphanageColl.Find(context.Background(), bson.D{})
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	defer func() {
		if err := orphanageCursor.Close(context.Background()); err != nil {
			panic(err)
		}
	}()
	// Проход по результатам и фильтрация по типу необходимости
	var result []interface{}
	for orphanageCursor.Next(context.Background()) {
		var orphanage map[string]interface{}
		if err := orphanageCursor.Decode(&orphanage); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//в переменную преобразуем тип интерфейс в примитив.обжектАйди
		str := orphanage["_id"].(primitive.ObjectID)

		// Фильтрация по типу необходимости
		needCursor, err := needsColl.Find(context.Background(), bson.M{"orphanageid": str.Hex(), "categoryofdonate": orphanageFilter.CategoryOfDonate})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Проходится по необходимостям если они есть и выводит
		for needCursor.Next(context.Background()) {
			result = append(result, orphanage["name"])
			var need map[string]interface{}
			if err := needCursor.Decode(&need); err != nil {
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
