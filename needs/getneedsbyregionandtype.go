package needs

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func GetNeedsByRegionAndType(w http.ResponseWriter, r *http.Request) {
	var err error
	err = flag.Set("logtostderr", "false") // Логировать в stderr (консоль) (false для записи в файл)
	if err != nil {
		log.Fatal(err)
	}
	err = flag.Set("stderrthreshold", "FATAL") // Устанавливаем порог для вывода ошибок в stderr
	if err != nil {
		log.Fatal(err)
	}
	err = flag.Set("log_dir", "C:/golang/logs/") // Указываем директорию для сохранения логов
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()
	defer glog.Flush()

	client := dbconnect.ConnectToDB()
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			glog.Fatal(err)
		}
	}()
	orphanageColl := client.Database("orphanage").Collection("orphanage")
	needsColl := client.Database("orphanage").Collection("need")
	// Парсинг данных из тела запроса
	var needFilter structures.NeedFilter
	if err = json.NewDecoder(r.Body).Decode(&needFilter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	needCursor, err := needsColl.Find(context.Background(), bson.M{"categoryofdonate": needFilter.CategoryOfDonate})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	var result []interface{}
	for needCursor.Next(context.Background()) {
		var need map[string]interface{}
		if err = needCursor.Decode(&need); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		objId, err := primitive.ObjectIDFromHex(need["orphanageid"].(string))
		if err != nil {
			// Обработка ошибки при парсинге строки в ObjectID
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		//
		var orphanage map[string]interface{}
		err = orphanageColl.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&orphanage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		if orphanage["region"] == needFilter.Region {
			result = append(result, need)
		}
	}
	err = needCursor.Close(context.Background())
	if err != nil {
		glog.Fatal(err)
	}
	if result == nil {
		needCursor, err = needsColl.Find(context.Background(), bson.M{"categoryofdonate": needFilter.CategoryOfDonate})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		for needCursor.Next(context.Background()) {
			var need map[string]interface{}
			if err = needCursor.Decode(&need); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				glog.Fatal(err)
			}
			result = append(result, need)
		}
		err = needCursor.Close(context.Background())
		if err != nil {
			glog.Fatal(err)
		}
	}
	glog.Info("Success")

	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		glog.Fatal(err)
	}
}
