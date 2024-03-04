package schedule

import (
	"awesomeProject1/dbconnect"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func DeactivateNeedIfExpired() {
	var err error
	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	needColl := client.Database("orphanage").Collection("need")
	currentDay := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, time.Now().Nanosecond(), time.Now().Location())
	nextDay := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, time.Now().Nanosecond(), time.Now().Location()).AddDate(0, 0, 1)
	needCursor, err := needColl.Find(context.Background(), bson.M{"expiring": bson.M{"$gte": currentDay, "$lt": nextDay}})
	if err != nil {
		log.Fatal(err)
	}
	for needCursor.Next(context.Background()) {
		var need map[string]interface{}
		err = needCursor.Decode(&need)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		_, err = needColl.UpdateOne(context.Background(), bson.D{{"_id", need["_id"].(primitive.ObjectID)}}, bson.D{{"$set", bson.D{{"isactive", false}}}})
		if err != nil {
			log.Fatal(err)
		}
	}
}
