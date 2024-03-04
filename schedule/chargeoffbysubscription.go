package schedule

import (
	"awesomeProject1/dbconnect"
	"awesomeProject1/structures"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func ChargeOffBySubscription() {
	var err error
	client := dbconnect.ConnectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	donateSubscribeColl := client.Database("orphanage").Collection("donatesubscribe")
	currentDay := time.Now().Day()
	subscribeCursor, err := donateSubscribeColl.Find(context.Background(), bson.M{"whichday": int32(currentDay)})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println(err)
			return
		}
		log.Fatal(err)
	}
	bankDetailsColl := client.Database("orphanage").Collection("bankdetails")
	orphanageColl := client.Database("orphanage").Collection("orphanage")
	donationHistoryColl := client.Database("orphanage").Collection("donationhistory")
	for subscribeCursor.Next(context.Background()) {
		var donateInfo structures.DonateSubscribe
		err = subscribeCursor.Decode(&donateInfo)
		if err != nil {
			log.Fatal(err)
		}
		objId, err := primitive.ObjectIDFromHex(donateInfo.BankDetailsId)
		//ПРОВЕРКА НА ДОСТАТОК БАЛАНСА
		var bankDetail map[string]interface{}
		err = bankDetailsColl.FindOne(context.Background(), bson.D{{"_id", objId}}).Decode(&bankDetail)
		if bankDetail["bill"].(int64) >= int64(donateInfo.Amount) {
			if err != nil {
				log.Fatal(err)
			}
			_, err = bankDetailsColl.UpdateOne(context.Background(), bson.D{{"_id", objId}}, bson.D{{"$inc", bson.D{{"bill", -donateInfo.Amount}}}})
			orphanageCount := len(donateInfo.OrphanageId)
			for i := 0; i < orphanageCount; i++ {
				objId, err = primitive.ObjectIDFromHex(donateInfo.OrphanageId[i])
				if err != nil {
					log.Fatal(err)
				}
				_, err = orphanageColl.UpdateOne(context.Background(), bson.D{{"_id", objId}}, bson.D{{"$inc", bson.D{{"bill", donateInfo.Amount / int64(orphanageCount)}}}})
			}
			doc := structures.DonationHistory{
				UserId:      "null",
				OrphanageId: donateInfo.OrphanageId,
				Sum:         int(donateInfo.Amount),
				Date:        primitive.NewDateTimeFromTime(time.Now().UTC()),
			}
			_, err = donationHistoryColl.InsertOne(context.Background(), doc)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Println(bankDetail["_id"], " not enough money")
		}
	}
	err = subscribeCursor.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
