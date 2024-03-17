package auth

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
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
	// Парсинг данных из тела запроса
	var loginData structures.LoginData
	if err = json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	// collection connect
	userColl := client.Database("orphanage").Collection("users")
	adminsColl := client.Database("orphanage").Collection("admins")
	// Search data by phone and password
	var result bson.M
	err = userColl.FindOne(context.TODO(), bson.D{{"phone", loginData.Phone}, {"password", loginData.Password}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = adminsColl.FindOne(context.TODO(), bson.D{{"phone", loginData.Phone}, {"password", loginData.Password}}).Decode(&result)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					result = bson.M{"result": "No documents"}
				} else {
					glog.Fatal(err)
				}
			}
		}
	}
	glog.Info("Successfully signed in")
	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		glog.Fatal(err)
	}
}
