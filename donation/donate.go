package donation

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
	"time"
)

func AddDonate(w http.ResponseWriter, r *http.Request) {
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
	// CONNECT TO COLLECTIONS
	bankDetailsColl := client.Database("orphanage").Collection("bankdetails")
	orphanageColl := client.Database("orphanage").Collection("orphanage")
	donationHistoryColl := client.Database("orphanage").Collection("donationhistory")
	userColl := client.Database("orphanage").Collection("users")
	// GET DATA FROM REQUEST
	var donate structures.Donate
	if err = json.NewDecoder(r.Body).Decode(&donate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	// TURN ID TO ObjectID
	objId, err := primitive.ObjectIDFromHex(donate.BankDetailsId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	var userDetails map[string]interface{}
	// SEARCH BANK DETAIL
	err = bankDetailsColl.FindOne(context.Background(), bson.D{{"_id", objId}}).Decode(&userDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	// Check if money is enough
	if userDetails["bill"].(int64) >= int64(donate.Sum) {
		// Money withdraw from bank card
		_, err = bankDetailsColl.UpdateOne(context.Background(), bson.D{{"_id", objId}}, bson.D{{"$inc", bson.D{{"bill", -donate.Sum}}}})
		if err != nil {
			glog.Fatal(err)
		}
		objId, err = primitive.ObjectIDFromHex(userDetails["userid"].(string)) // get userid
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		// Increase donated by user
		_, err = userColl.UpdateOne(context.Background(), bson.D{{"_id", objId}}, bson.D{{"$inc", bson.D{{"donated", donate.Sum}}}})
		if err != nil {
			glog.Fatal(err)
		}
		// Count orphanages
		orphanageCount := len(donate.OrphanageId)
		// Increase orphanage's bill, if there are many orphanages then divide money to all
		for i := 0; i < orphanageCount; i++ {
			orphanageId, err := primitive.ObjectIDFromHex(donate.OrphanageId[i])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				glog.Fatal(err)
			}
			_, err = orphanageColl.UpdateOne(context.Background(), bson.D{{"_id", orphanageId}}, bson.D{{"$inc", bson.D{{"bill", donate.Sum / orphanageCount}}}})
			if err != nil {
				glog.Fatal(err)
			}
		}
		// Add data to donation history
		doc := structures.DonationHistory{
			UserId:      userDetails["userid"].(string),
			OrphanageId: donate.OrphanageId,
			Sum:         donate.Sum,
			Date:        primitive.NewDateTimeFromTime(time.Now().Add(5 * time.Hour)),
		}
		// Insert donation history
		_, err = donationHistoryColl.InsertOne(context.Background(), doc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		glog.Info("Added successfully")
	} else {
		// if in bank card is not enough money
		http.Error(w, "Недостаточно средств", http.StatusInternalServerError)
		glog.Fatal(err)
	}
	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Success")
	if err != nil {
		glog.Fatal(err)
	}
}
