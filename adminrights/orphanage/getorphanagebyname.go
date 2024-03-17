package orphanage

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func GetOrphanageByName(w http.ResponseWriter, r *http.Request) {
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
		glog.Fatal(err)
	}
	// Подключение к базе данных
	client := dbconnect.ConnectToDB()
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			glog.Fatal(err)
		}
	}()
	coll := client.Database("orphanage").Collection("orphanage") // Collection connect
	// Получение ID коллекции из URL
	name := r.URL.Query().Get("name")
	if name == "" {
		// ЕСЛИ НЕТ ИМЕНИ ТО ВЫВОДИТ ВСЕ ОРФЕНЕЙДЖИ
		cursor, err := coll.Find(context.Background(), bson.D{})
		if err != nil {
			fmt.Println(err)
			glog.Fatal(err)
		}
		var result []interface{}
		for cursor.Next(context.Background()) {
			var cur map[string]interface{}
			if err = cursor.Decode(&cur); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				glog.Fatal(err)
			}
			result = append(result, cur)
		}
		// Отправка данных в формате JSON
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(result)
		if err = cursor.Close(context.Background()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		return
	}
	// Поиск данных по имени
	var orphanage structures.Orphanage
	err = coll.FindOne(context.Background(), bson.M{"name": name}).Decode(&orphanage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	glog.Info("Success")
	// Отправка данных в формате JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	err = json.NewEncoder(w).Encode(orphanage)
	if err != nil {
		glog.Fatal(err)
	}
}
