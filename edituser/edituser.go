package edituser

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func EditUser(w http.ResponseWriter, r *http.Request) {
	var err error

	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var userData structures.Users
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	coll := client.Database("orphanage").Collection("users")

	filter := bson.D{{"_id", userData.ID}}
	update := bson.D{{"$set", userData}}

	_, err = coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Successfully edited")
	if err != nil {
		return
	}
}
