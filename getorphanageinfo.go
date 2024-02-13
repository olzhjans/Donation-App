package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type Orphanage struct {
	//ID           primitive.ObjectID `bson:"_id"`
	Name         string
	Region       string
	Address      string
	Description  string
	ChildsCount  string
	WorkingHours string
}

func getOrphanageInfo(orphanageName string) {
	client := connectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("orphanage").Collection("orphanage")

	var filter bson.D
	if orphanageName == "" {
		filter = bson.D{}
	} else {
		filter = bson.D{{"name", orphanageName}}
	}
	// Retrieves documents that match the query filer
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	// end find
	var results []Orphanage
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	// Prints the results of the find operation as structs
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
}
