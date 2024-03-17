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

func GetDonationHistoryByUserId(w http.ResponseWriter, r *http.Request) {
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
	// Get userid from URL
	userid := r.URL.Query().Get("userid")
	// Search user's donation history
	donationCursor, err := donationHistoryColl.Find(context.Background(), bson.M{"user-id": userid})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	// RETURN HISTORY
	var result []interface{}
	for donationCursor.Next(context.Background()) {
		var donate structures.DonationHistory
		err = donationCursor.Decode(&donate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		result = append(result, donate)
	}
	// Close cursor
	err = donationCursor.Close(context.Background())
	if err != nil {
		glog.Error(err)
	}
	glog.Info("Success")
	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		glog.Error(err)
	}
}
