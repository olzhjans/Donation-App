package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func signIn() bool {
	client := connectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	//SIGN IN AND PRINT DATA
	enteredPhone, enteredPassword := getSignInData()
	userColl := client.Database("orphanage").Collection("users")
	adminsColl := client.Database("orphanage").Collection("admins")
	var result bson.M
	err := userColl.FindOne(context.TODO(), bson.D{{"phone", enteredPhone}, {"password", enteredPassword}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = adminsColl.FindOne(context.TODO(), bson.D{{"phone", enteredPhone}, {"password", enteredPassword}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document was found with this telephone number and password")
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
	if result["who"] == "Moderator" {
		return true
	}
	return false
}

func getSignInData() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	phone := ""

	fmt.Println("SIGN IN")
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
