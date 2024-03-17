package comments

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
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
	// Collection connect
	commentaryColl := client.Database("orphanage").Collection("comments")
	// GET DATA FROM REQUEST
	var commentary structures.Commentary
	if err = json.NewDecoder(r.Body).Decode(&commentary); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Fatal(err)
	}
	// Add actual datetime
	commentary.Date = primitive.NewDateTimeFromTime(time.Now().Add(5 * time.Hour))
	// Inserting data
	insertedComment, err := commentaryColl.InsertOne(context.Background(), commentary)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		glog.Fatal(err)
	}
	// COUNT COMMENTS, IF > 100 THEN DELETE EXTRA
	count, err := commentaryColl.CountDocuments(context.Background(), bson.D{{"need-id", commentary.NeedId}})
	if err != nil {
		log.Fatal(err)
	}
	if count > 100 {
		numToDelete := count - 100
		// Находим самые старые лишние записи
		cursor, err := commentaryColl.Find(context.Background(), bson.D{{"need-id", commentary.NeedId}}, options.Find().SetSort(map[string]interface{}{"date": 1}).SetLimit(int64(numToDelete)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		// Проходимся по старым записям
		for cursor.Next(context.Background()) {
			var record structures.Commentary
			if err = cursor.Decode(&record); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				glog.Fatal(err)
			}
			// Удаляем запись
			_, err = commentaryColl.DeleteOne(context.Background(), bson.M{"_id": record.ID})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				glog.Fatal(err)
			}
			glog.Info("Удалена старая запись с ID: %s\n", record.ID)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
		// Cursor close
		err = cursor.Close(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			glog.Fatal(err)
		}
	}
	glog.Info(insertedComment.InsertedID, "added successfully")
	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Added successfully")
	if err != nil {
		glog.Fatal(err)
	}
}
