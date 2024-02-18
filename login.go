package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func logIn() (string, bool) {
	client := connectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	isAdmin := false

	//SIGN IN AND PRINT DATA
	enteredPhone, enteredPassword := getLogInData()
	userColl := client.Database("orphanage").Collection("users")
	adminsColl := client.Database("orphanage").Collection("admins")
	var result bson.M
	err := userColl.FindOne(context.TODO(), bson.D{{"phone", enteredPhone}, {"password", enteredPassword}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = adminsColl.FindOne(context.TODO(), bson.D{{"phone", enteredPhone}, {"password", enteredPassword}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document was found with this telephone number and password")
		} else {
			isAdmin = true
		}
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

	userObjectID := result["_id"]
	stringUserObjectID := userObjectID.(primitive.ObjectID).Hex()
	return stringUserObjectID, isAdmin
}

func getLogInData() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	phone := ""

	fmt.Println("LOG IN")
	for {
		fmt.Printf("Enter telephone number: ")
		scanner.Scan()
		phone = scanner.Text()
		if isValidPhoneNumber(phone) == true {
			break
		} else {
			fmt.Println("Incorrect phone number")
		}
	}

	fmt.Printf("Enter password: ")
	scanner.Scan()
	password := scanner.Text()

	return phone, password
}