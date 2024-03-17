package orphanage

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func EditOrphanage(w http.ResponseWriter, r *http.Request) {
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
	// CONNECT TO DB
	client := dbconnect.ConnectToDB()
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			glog.Fatal(err)
		}
	}()
	// GET DATA FROM REQUEST
	var orphanageData structures.Orphanage
	if err = json.NewDecoder(r.Body).Decode(&orphanageData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	// IMPLEMENT EMPTY "updateFields"
	updateFields := bson.D{}
	// CHECK IF NOT EMPTY THEN ADD TO "updateFields"
	if orphanageData.Name != "" {
		updateFields = append(updateFields, bson.E{Key: "name", Value: orphanageData.Name})
	}
	if orphanageData.Region != "" {
		updateFields = append(updateFields, bson.E{Key: "region", Value: orphanageData.Region})
	}
	if orphanageData.Address != "" {
		updateFields = append(updateFields, bson.E{Key: "address", Value: orphanageData.Address})
	}
	if orphanageData.Description != "" {
		updateFields = append(updateFields, bson.E{Key: "description", Value: orphanageData.Description})
	}
	if orphanageData.ChildsCount != 0 {
		updateFields = append(updateFields, bson.E{Key: "childs-count", Value: orphanageData.ChildsCount})
	}
	if orphanageData.WorkingHours != "" {
		updateFields = append(updateFields, bson.E{Key: "working-hours", Value: orphanageData.WorkingHours})
	}
	if orphanageData.Photos != nil {
		updateFields = append(updateFields, bson.E{Key: "photos", Value: orphanageData.Photos})
	}
	if orphanageData.Bill != 0 {
		updateFields = append(updateFields, bson.E{Key: "bill", Value: orphanageData.Bill})
	}
	var result interface{}
	var update bson.D
	// RESPONSE
	if len(updateFields) > 0 {
		update = bson.D{{"$set", updateFields}}
	} else {
		// If request is empty then response "Not Found"
		result = "No data typed"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			glog.Fatal(err)
		}
		glog.Error("No data typed")
		return
	}
	// Connect to collection
	coll := client.Database("orphanage").Collection("orphanage")
	filter := bson.D{{"_id", orphanageData.ID}}
	// Updating data
	updated, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		glog.Fatal(err)
	}
	// RESPONSE
	if updated.MatchedCount == 0 {
		glog.Info(orphanageData.ID, " not found")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		result = orphanageData.ID.Hex() + " not found"
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			glog.Fatal(err)
		}
		return
	}
	glog.Info(orphanageData.ID, " edited successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	result = orphanageData.ID.Hex() + " successfully edited"
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		glog.Fatal(err)
	}
}
