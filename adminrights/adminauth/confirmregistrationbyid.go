package adminauth

import (
	"awesomeProject1/adminrights/mail"
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

func ConfirmRegistrationById(w http.ResponseWriter, r *http.Request) {
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
	// Проверка метода запроса
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		glog.Fatal("http.StatusMethodNotAllowed")
	}
	// Подключение к базе данных
	client := dbconnect.ConnectToDB()
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			glog.Fatal(err)
		}
	}()
	waitingColl := client.Database("orphanage").Collection("waitinglist")
	// Получение ID коллекции из URL
	id := r.URL.Query().Get("_id")
	objID, err := primitive.ObjectIDFromHex(id) // TURNING id to ObjectID
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	// SEARCH DATA BY "_id"
	var admin structures.Admins
	err = waitingColl.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	// WRITE ACTUAL DATETIME
	admin.SignupDate = primitive.NewDateTimeFromTime(time.Now().Add(5 * time.Hour))
	// Connecting to collection
	adminColl := client.Database("orphanage").Collection("admins")
	// Inserting data to collection
	_, err = adminColl.InsertOne(context.Background(), admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	// DELETE FROM WAITING LIST
	_, err = waitingColl.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	err = json.NewEncoder(w).Encode("Confirmed successfully")
	if err != nil {
		glog.Fatal(err)
	}
	glog.Info(objID, "confirmed successfully")
	// SEND MAIL
	mail.SendMail("olzhjans@gmail.com", "Donation-App", "Congratulations, your request has been confirmed")
}
