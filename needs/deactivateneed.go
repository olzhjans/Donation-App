package needs

import (
	"awesomeProject1/dbconnect"
	"context"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func DeactivateNeedByNeedId(w http.ResponseWriter, r *http.Request) {
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
	coll := client.Database("orphanage").Collection("need")
	// Получение ID коллекции из URL
	needId := r.URL.Query().Get("needid")
	objId, err := primitive.ObjectIDFromHex(needId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	_, err = coll.UpdateOne(context.Background(), bson.D{{"_id", objId}}, bson.D{{"$set", bson.D{{"isactive", false}}}})
	if err != nil {
		glog.Fatal(err)
	}
	glog.Info(objId, " deactivated successfully")

	// Отправка данных в формате JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Successfully deactivated")
	if err != nil {
		glog.Fatal(err)
	}
}
