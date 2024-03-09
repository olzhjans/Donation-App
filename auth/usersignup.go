package auth

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/mail"
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

func UserSignUp(w http.ResponseWriter, r *http.Request) {
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
	coll := client.Database("orphanage").Collection("users")
	// Парсинг данных из тела запроса
	var user structures.Users
	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	user.Donated = 0
	user.SignupDate = primitive.NewDateTimeFromTime(time.Now().Add(5 * time.Hour))
	// Вставка данных в базу данных
	insertedUser, err := coll.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	glog.Info(insertedUser.InsertedID, "successfully signed up")

	w.Header().Set("Content-Type", "application/json")
	// Возвращаем успешный статус
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Added successfully")
	if err != nil {
		glog.Fatal(err)
	}
	//SEND MAIL
	mail.SendMail("olzhjans@gmail.com", "Donation-App", "Congratulations! You have registered")
}
