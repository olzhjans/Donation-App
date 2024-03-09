package needs

import (
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

func AddNeeds(w http.ResponseWriter, r *http.Request) {
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

	coll := client.Database("orphanage").Collection("need")
	// Парсинг данных из тела запроса
	var need structures.Need
	if err = json.NewDecoder(r.Body).Decode(&need); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	need.Expiring = primitive.NewDateTimeFromTime(time.Now().UTC().Add(5*time.Hour).AddDate(0, 1, 0))
	// Вставка данных в базу данных
	insertedNeed, err := coll.InsertOne(context.Background(), need)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	glog.Info(insertedNeed.InsertedID, " added successfully")
	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Successfully added")
	if err != nil {
		glog.Fatal(err)
	}
}
