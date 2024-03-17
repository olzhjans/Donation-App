package wherespent

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

func ShowWhereSpent(w http.ResponseWriter, r *http.Request) {
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
	// Подключение к базе данных
	client := dbconnect.ConnectToDB()
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			glog.Fatal(err)
		}
	}()
	// Collection connect
	coll := client.Database("orphanage").Collection("wherespent")
	// GET DATA FROM REQUEST
	var whereSpentFilter structures.WhereSpentFilter
	if err = json.NewDecoder(r.Body).Decode(&whereSpentFilter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	// SEARCH
	cursor, err := coll.Find(context.Background(), bson.M{"orphanageid": whereSpentFilter.OrphanageId, "date": bson.M{"$gte": whereSpentFilter.From, "$lte": whereSpentFilter.To}})
	if err != nil {
		panic(err)
	}
	// DECODE
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
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(result)
}
