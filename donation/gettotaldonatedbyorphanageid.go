package donation

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

func GetTotalDonatedByOrphanageIdAndPeriod(w http.ResponseWriter, r *http.Request) {
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
	// Collection connect
	donationHistoryColl := client.Database("orphanage").Collection("donationhistory")
	// Get data from request
	var donation structures.DonationFilter
	if err = json.NewDecoder(r.Body).Decode(&donation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	// Search history
	donationCursor, err := donationHistoryColl.Find(context.Background(), bson.M{"orphanage-id": donation.Id, "date": bson.M{"$gte": donation.From, "$lte": donation.To}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	// SUM ALL DONATES
	sum := 0
	for donationCursor.Next(context.Background()) {
		var donate structures.DonationHistory
		err = donationCursor.Decode(&donate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		sum += donate.Sum / len(donate.OrphanageId)
	}
	// cursor close
	err = donationCursor.Close(context.Background())
	if err != nil {
		glog.Fatal(err)
	}
	glog.Info("Success")
	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(sum)
	if err != nil {
		glog.Fatal(err)
	}
}
