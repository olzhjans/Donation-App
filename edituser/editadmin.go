package edituser

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

func EditAdmin(w http.ResponseWriter, r *http.Request) {
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

	var adminData structures.Admins
	if err = json.NewDecoder(r.Body).Decode(&adminData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	updateFields := bson.D{}
	if adminData.Phone != "" {
		updateFields = append(updateFields, bson.E{Key: "phone", Value: adminData.Phone})
	}
	if adminData.Password != "" {
		updateFields = append(updateFields, bson.E{Key: "password", Value: adminData.Password})
	}
	if adminData.Firstname != "" {
		updateFields = append(updateFields, bson.E{Key: "firstname", Value: adminData.Firstname})
	}
	if adminData.Lastname != "" {
		updateFields = append(updateFields, bson.E{Key: "lastname", Value: adminData.Lastname})
	}
	if adminData.Email != "" {
		updateFields = append(updateFields, bson.E{Key: "email", Value: adminData.Email})
	}
	if adminData.Region != "" {
		updateFields = append(updateFields, bson.E{Key: "region", Value: adminData.Region})
	}
	if adminData.Who != "" {
		updateFields = append(updateFields, bson.E{Key: "who", Value: adminData.Who})
	}
	if adminData.Id != "" {
		updateFields = append(updateFields, bson.E{Key: "id", Value: adminData.Id})
	}
	if adminData.SignupDate != 0 {
		updateFields = append(updateFields, bson.E{Key: "signupdate", Value: adminData.SignupDate})
	}
	if adminData.OrphanageId != "" {
		updateFields = append(updateFields, bson.E{Key: "orphanageid", Value: adminData.OrphanageId})
	}
	var result interface{}
	var update bson.D
	if len(updateFields) > 0 {
		update = bson.D{{"$set", updateFields}}
	} else {
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
	filter := bson.D{{"_id", adminData.ID}}
	coll := client.Database("orphanage").Collection("admins")
	updated, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		glog.Fatal(err)
	}
	if updated.MatchedCount == 0 {
		glog.Info(adminData.ID, " not found")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		result = adminData.ID.Hex() + " not found"
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			glog.Fatal(err)
		}
		return
	}
	glog.Info(adminData.ID, " edited successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	result = adminData.ID.Hex() + " successfully edited"
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		glog.Fatal(err)
	}
}
