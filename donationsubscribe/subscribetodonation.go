package donationsubscribe

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func SubscribeToDonation(w http.ResponseWriter, r *http.Request) {
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
			panic(err)
		}
	}()
	// Collection connect
	bankDetailsColl := client.Database("orphanage").Collection("bankdetails")
	donateSubscribeColl := client.Database("orphanage").Collection("donatesubscribe")
	// Get data from request
	var subscribe structures.DonationSubscribe
	if err = json.NewDecoder(r.Body).Decode(&subscribe); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	// CHECK IF BANK CARD IS EXIST
	var bankDetails map[string]interface{}
	err = bankDetailsColl.FindOne(context.Background(), bson.D{{"cardnumber", subscribe.BankDetails.CardNumber}}).Decode(&bankDetails)
	var bankCardId interface{}
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// IF CARD NOT FOUND THEN ADD
			result, err := bankDetailsColl.InsertOne(context.Background(), subscribe.BankDetails)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				glog.Fatal(err)
			}
			bankCardId = result.InsertedID
			fmt.Println("Bank card saved!")
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
	} else {
		// IF BANK CARD IS EXIST
		bankCardId = bankDetails["_id"]
		fmt.Println("Bank card exist.")
	}
	// ADD SUBSCRIPTION TO COLLECTION
	inserted, err := donateSubscribeColl.InsertOne(context.Background(), bson.D{{"orphanageid", subscribe.OrphanageId}, {"bankdetailsid", bankCardId.(primitive.ObjectID).Hex()}, {"amount", subscribe.Amount}, {"whichday", subscribe.WhichDay}, {"isactive", subscribe.IsActive}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	glog.Info(inserted.InsertedID, " added subscribe successfully")
	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Success")
	if err != nil {
		glog.Fatal(err)
	}
}
