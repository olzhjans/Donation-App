package adminauth

import (
	"awesomeProject1/adminrights/mail"
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

func AdminSignUp(w http.ResponseWriter, r *http.Request) {
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
	// DB CONNECTION
	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			glog.Fatal(err)
		}
	}()
	coll := client.Database("orphanage").Collection("waitinglist")
	// Парсинг данных из тела запроса
	var admin structures.Admins
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	// ADDING ACTUAL DATETIME TO STRUCT
	admin.SignupDate = primitive.NewDateTimeFromTime(time.Now().Add(5 * time.Hour))
	// INSERTING DATA TO COLLECTION
	insertedAdmin, err := coll.InsertOne(context.Background(), admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	glog.Info(insertedAdmin.InsertedID, "successfully signed up")
	//	RETURN RESPONSE
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Added successfully")
	if err != nil {
		glog.Fatal(err)
	}
	//	SEND MAIL
	mail.SendMail("olzhjans@gmail.com", "Donation-App", "Your registration request has been sent.")
}
