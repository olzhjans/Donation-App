package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func editUserData(userID string, isAdmin bool) {
	var err error

	client := connectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		panic(err)
	}

	coll := client.Database("orphanage").Collection("users")
	if isAdmin {
		coll = client.Database("orphanage").Collection("admins")
	}

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

	password := getUserData()

	filter := bson.D{{"_id", result["_id"]}}

	update := bson.D{{"$set", bson.D{{"password", password}}}}

	opts := options.Update().SetUpsert(true)

	_, err = coll.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		panic(err)
	}

	fmt.Println("SUCCESS")
}

func getUserData() string {
	scanner := bufio.NewScanner(os.Stdin)
	var password string
	fmt.Println("Type new password")
	for {
		fmt.Printf("New password: ")
		scanner.Scan()
		password = scanner.Text()
		if isValidPassword(password) {
			break
		}
	}
	return password
}
