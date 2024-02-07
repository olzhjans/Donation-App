package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func signIn() {
	//CONNECTING TO MONGODB
	if err := godotenv.Load("mongodb.env"); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
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
	err = userColl.FindOne(context.TODO(), bson.D{{"phone", enteredPhone}, {"password", enteredPassword}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = adminsColl.FindOne(context.TODO(), bson.D{{"phone", enteredPhone}, {"password", enteredPassword}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found with this telephone number and password")
		}
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
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
	//fmt.Println(enteredPassword)

	return phone, password
}
