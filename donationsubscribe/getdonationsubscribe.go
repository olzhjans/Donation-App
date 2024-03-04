package donationsubscribe

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

func GetDonationSubscribeByUserId(w http.ResponseWriter, r *http.Request) {
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
	bankDetailsColl := client.Database("orphanage").Collection("bankdetails")
	donateSubscribeColl := client.Database("orphanage").Collection("donatesubscribe")
	userid := r.URL.Query().Get("userid")
	bankDetailsCursor, err := bankDetailsColl.Find(context.Background(), bson.D{{"userid", userid}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	var result []interface{}
	for bankDetailsCursor.Next(context.Background()) {
		var detail structures.BankDetails
		if err = bankDetailsCursor.Decode(&detail); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		donateSubscriptionCursor, err := donateSubscribeColl.Find(context.Background(), bson.D{{"bankdetailsid", detail.ID.Hex()}})
		for donateSubscriptionCursor.Next(context.Background()) {
			var subscription map[string]interface{}
			if err = donateSubscriptionCursor.Decode(&subscription); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				glog.Fatal(err)
			}
			result = append(result, subscription)
		}
		err = donateSubscriptionCursor.Close(context.Background())
		if err != nil {
			glog.Fatal(err)
		}
	}
	err = bankDetailsCursor.Close(context.Background())
	if err != nil {
		glog.Fatal(err)
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
