package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func editOrphanageData(orphanageName string) {
	client := connectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("orphanage").Collection("orphanage")

	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"name", orphanageName}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No orphanage was found with ", orphanageName, " name")
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

	name, region, address, description, childsCount, workingHours := getOrphanageData()

	filter := bson.D{{"_id", result["_id"]}}
	//fmt.Println(filter)

	doc := Orphanage{
		Name:         name,
		Region:       region,
		Address:      address,
		Description:  description,
		ChildsCount:  childsCount,
		WorkingHours: workingHours,
	}

	update := bson.D{{"$set", doc}}

	opts := options.Update().SetUpsert(true)

	_, err = coll.UpdateOne(context.TODO(), filter, update, opts)
	//_, err := coll.ReplaceOne(context.TODO(), filter, doc)
	if err != nil {
		panic(err)
	}

	fmt.Println("SUCCESS")
}

func getOrphanageData() (string, string, string, string, string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	var name, region, address, description, childsCount, workingHours string
	fmt.Println("Type information to edit orphanage")
	for {
		fmt.Printf("Enter orphanage name: ")
		scanner.Scan()
		name = scanner.Text()
		if name != "" {
			break
		}
	}
	for {
		fmt.Printf("Enter orphanage region: ")
		scanner.Scan()
		region = scanner.Text()
		if region != "" {
			break
		}
	}
	for {
		fmt.Printf("Enter orphanage address: ")
		scanner.Scan()
		address = scanner.Text()
		if address != "" {
			break
		}
	}
	for {
		fmt.Printf("Enter orphanage description: ")
		scanner.Scan()
		description = scanner.Text()
		if description != "" {
			break
		}
	}
	for {
		fmt.Printf("Enter orphanage childs count: ")
		scanner.Scan()
		childsCount = scanner.Text()
		if childsCount != "" {
			break
		}
	}
	for {
		fmt.Printf("Enter orphanage working hours: ")
		scanner.Scan()
		workingHours = scanner.Text()
		if workingHours != "" {
			break
		}
	}

	return name, region, address, description, childsCount, workingHours
}
