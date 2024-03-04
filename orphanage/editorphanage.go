package orphanage

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

func EditOrphanage(w http.ResponseWriter, r *http.Request) {
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

	// Парсинг данных из тела запроса
	var orphanageData structures.Orphanage
	if err = json.NewDecoder(r.Body).Decode(&orphanageData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	coll := client.Database("orphanage").Collection("orphanage")
	filter := bson.D{{"_id", orphanageData.ID}}
	update := bson.D{{"$set", orphanageData}} //изменить все поля, если нет какого-то поля то добавить
	_, err = coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		glog.Fatal(err)
	}
	glog.Info(orphanageData.ID, " edited successfully")

	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Successfully edited")
	if err != nil {
		glog.Fatal(err)
	}
}
