package comments

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

func GetComments(w http.ResponseWriter, r *http.Request) {
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
		if err := client.Disconnect(context.TODO()); err != nil {
			glog.Fatal(err)
		}
	}()
	// Collection connect
	commentaryColl := client.Database("orphanage").Collection("comments")
	// GET DATA FROM REQUEST
	var commentsFilter structures.CommentaryFilter
	if err = json.NewDecoder(r.Body).Decode(&commentsFilter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	// Search data and write to cursor
	commentaryCursor, err := commentaryColl.Find(context.Background(), bson.M{"need-id": commentsFilter.NeedId, "date": bson.M{"$gte": commentsFilter.From, "$lte": commentsFilter.To}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	var result []interface{}
	// Проходимся по курсору
	for commentaryCursor.Next(context.Background()) {
		var commentary map[string]interface{}
		if err = commentaryCursor.Decode(&commentary); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		result = append(result, commentary)
	}
	// Cursor close
	err = commentaryCursor.Close(context.Background())
	if err != nil {
		glog.Fatal(err)
	}
	glog.Info("Success")
	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		glog.Fatal(err)
	}
}
