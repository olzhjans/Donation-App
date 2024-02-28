package comments

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func GetComments(w http.ResponseWriter, r *http.Request) {
	var err error
	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	commentaryColl := client.Database("orphanage").Collection("comments")
	var CommentsFilter structures.CommentaryFilter
	if err := json.NewDecoder(r.Body).Decode(&CommentsFilter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	commentaryCursor, err := commentaryColl.Find(context.Background(), bson.M{"needid": CommentsFilter.NeedId, "date": bson.M{"$gte": CommentsFilter.From, "$lte": CommentsFilter.To}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var result []interface{}
	for commentaryCursor.Next(context.Background()) {
		var commentary map[string]interface{}
		if err = commentaryCursor.Decode(&commentary); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result = append(result, commentary)
	}
	err = commentaryCursor.Close(context.Background())
	if err != nil {
		return
	}
	// Возвращаем успешный статус
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}
