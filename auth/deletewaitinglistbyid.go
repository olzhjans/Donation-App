package auth

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/mail"
	"context"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func DeleteWaitingListById(w http.ResponseWriter, r *http.Request) {
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
	if r.Method != http.MethodDelete {
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
	coll := client.Database("orphanage").Collection("waitinglist")
	// Получение ID записи из запроса
	params := r.URL.Query()
	id := params.Get("id")
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		glog.Fatal("ID parameter is required, http.StatusBadRequest")
	}
	// Преобразование ID в формат BSON
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	filter := bson.M{"_id": objID}
	_, err = coll.DeleteOne(context.Background(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	glog.Info(objID, "deleted successfully")
	// Отправка ответа об успешном удалении
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Deleted successfully")
	if err != nil {
		glog.Fatal(err)
	}
	//SEND MAIL
	mail.SendMail("olzhjans@gmail.com", "Donation-App", "Unfortunately, your request has been rejected")
}
