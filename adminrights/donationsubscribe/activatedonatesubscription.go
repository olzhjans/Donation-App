package donationsubscribe

import (
	"awesomeProject1/dbconnect"
	"context"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func ActivateDonateSubscription(w http.ResponseWriter, r *http.Request) {
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
	// DB connect
	client := dbconnect.ConnectToDB()
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			glog.Fatal(err)
		}
	}()
	donateSubscribeColl := client.Database("orphanage").Collection("donatesubscribe") // Collection connect
	// GET ID FROM URL
	id := r.URL.Query().Get("_id")
	// TURN "_id" TO OBJECTID
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	// UPDATING DATA
	_, err = donateSubscribeColl.UpdateOne(context.Background(), bson.D{{"_id", objId}}, bson.D{{"$set", bson.D{{"isactive", true}}}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	glog.Info(id, "activated successfully")
	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode("Successfully activated")
	if err != nil {
		glog.Fatal(err)
	}
}
