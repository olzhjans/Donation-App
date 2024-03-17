package needs

import (
	"awesomeProject1/dbconnect"
	"context"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func ShowNeeds(w http.ResponseWriter, r *http.Request) {
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
	// Collection connect
	coll := client.Database("orphanage").Collection("need")
	// Получение ID коллекции из URL
	orphanageid := r.URL.Query().Get("orphanageid")
	// Search
	cursor, err := coll.Find(context.Background(), bson.M{"orphanageid": orphanageid})
	if err != nil {
		glog.Fatal(err)
	}
	// If there is no id in URL then write all needs
	if orphanageid == "" {
		cursor, err = coll.Find(context.Background(), bson.D{})
		if err != nil {
			glog.Fatal(err)
		}
	}
	// Decode needs
	var result []interface{}
	for cursor.Next(context.Background()) {
		var cur map[string]interface{}
		if err = cursor.Decode(&cur); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		result = append(result, cur)
	}
	if err = cursor.Close(context.Background()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	glog.Info("Success")
	// Отправка данных в формате JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	err = json.NewEncoder(w).Encode(result)
}
