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

func EditUser(w http.ResponseWriter, r *http.Request) {
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
	// DB CONNECT
	client := dbconnect.ConnectToDB()
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			glog.Fatal(err)
		}
	}()
	// GET DATA FROM REQUEST
	var userData structures.Users
	if err = json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	// IF DATA IS TYPED THEN ADD TO "updateFields"
	updateFields := bson.D{}
	if userData.Phone != "" {
		updateFields = append(updateFields, bson.E{Key: "phone", Value: userData.Phone})
	}
	if userData.Password != "" {
		updateFields = append(updateFields, bson.E{Key: "password", Value: userData.Password})
	}
	if userData.Firstname != "" {
		updateFields = append(updateFields, bson.E{Key: "firstname", Value: userData.Firstname})
	}
	if userData.Lastname != "" {
		updateFields = append(updateFields, bson.E{Key: "lastname", Value: userData.Lastname})
	}
	if userData.Email != "" {
		updateFields = append(updateFields, bson.E{Key: "email", Value: userData.Email})
	}
	if userData.Region != "" {
		updateFields = append(updateFields, bson.E{Key: "region", Value: userData.Region})
	}
	if userData.Donated != 0 {
		updateFields = append(updateFields, bson.E{Key: "donated", Value: userData.Donated})
	}
	if userData.SignupDate != 0 {
		updateFields = append(updateFields, bson.E{Key: "signupdate", Value: userData.SignupDate})
	}
	var result interface{}
	var update bson.D
	// Check if "updateFields" is not empty
	if len(updateFields) > 0 {
		update = bson.D{{"$set", updateFields}}
	} else {
		// RESPONSE IF EMPTY
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
	filter := bson.D{{"_id", userData.ID}}                   // implement filter
	coll := client.Database("orphanage").Collection("users") // collection connect
	// Update data
	updated, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		glog.Fatal(err)
	}
	if updated.MatchedCount == 0 {
		// Response if user not found
		glog.Info(userData.ID, " not found")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		result = userData.ID.Hex() + " not found"
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			glog.Fatal(err)
		}
		return
	}
	// RESPONSE IF USER FOUND
	glog.Info(userData.ID, " edited successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	result = userData.ID.Hex() + " successfully edited"
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		glog.Fatal(err)
	}
}
