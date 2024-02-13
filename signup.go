package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

type Users struct {
	Password   string
	Email      string
	Region     string
	Firstname  string
	Lastname   string
	Phone      string
	Donated    string
	SignupDate string
}

type Admins struct {
	Password    string
	Email       string
	Region      string
	Firstname   string
	Lastname    string
	Phone       string
	Who         string
	Id          string
	SignupDate  string
	OrphanageId string
}

func signUp() {
	client := connectToDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	//GET REGISTRATION INFO AND ADD TO DATABASE
	userType, firstName, lastName, phone, email, region, id, orphanageID, password := getSignUpData()
	if orphanageID == "" {
	} //убрать
	if userType == "1" { //USER
		userColl := client.Database("orphanage").Collection("users")
		currentTime := time.Now()
		doc := Users{
			Firstname:  firstName,
			Lastname:   lastName,
			Phone:      phone,
			Password:   password,
			Email:      email,
			Region:     region,
			Donated:    "0",
			SignupDate: currentTime.Format("02.01.2006")}

		result, _ := userColl.InsertOne(context.TODO(), doc)
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	} else { //MODERATOR
		adminColl := client.Database("orphanage").Collection("admins")
		currentTime := time.Now()
		doc := Admins{
			Firstname:   firstName,
			Lastname:    lastName,
			Phone:       phone,
			Password:    password,
			Email:       email,
			Region:      region,
			Who:         "Moderator",
			Id:          id,
			OrphanageId: orphanageID,
			SignupDate:  currentTime.Format("02.01.2006")}
		fmt.Println(doc)
		result, _ := adminColl.InsertOne(context.TODO(), doc)
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}

}

func getSignUpData() (string, string, string, string, string, string, string, string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	userType, firstName, lastName, phone, email, region, id, orphanageID, password := "", "", "", "", "", "", "", "", ""

	fmt.Println("SIGN UP")
	for {
		fmt.Println("User - 1, moderator - 2")
		fmt.Printf("Type: ")
		scanner.Scan()
		userType = scanner.Text()
		if userType == "1" || userType == "2" {
			break
		}
	}

	for {
		fmt.Printf("First name (Max 14 char): ")
		scanner.Scan()
		firstName = scanner.Text()
		if len(firstName) > 1 && len(firstName) <= 14 {
			break
		} else {
			fmt.Println("Incorrect length")
		}
	}

	for {
		fmt.Printf("Last name (Max 14 char): ")
		scanner.Scan()
		lastName = scanner.Text()
		if len(lastName) > 1 && len(lastName) <= 32 {
			break
		} else {
			fmt.Println("Incorrect length")
		}
	}

	for {
		fmt.Printf("Phone number (+7********** or 8**********): ")
		scanner.Scan()
		phone = scanner.Text()
		if isValidPhoneNumber(phone) {
			break
		} else {
			fmt.Println("Phone number is incorrect")
		}
	}

	for {
		fmt.Printf("E-mail (Max 32 char): ")
		scanner.Scan()
		email = scanner.Text()
		if len(email) <= 32 && isEmailValid(email) == true {
			break
		} else {
			fmt.Println("Email is incorrect")
		}
	}

	for {
		fmt.Printf("Region (Max 16 char): ")
		scanner.Scan()
		region = scanner.Text()
		if len(region) > 2 && len(region) <= 16 {
			break
		} else {
			fmt.Println("Incorrect length")
		}
	}

	for {
		fmt.Printf("Password (Min 8 char): ")
		scanner.Scan()
		password = scanner.Text()
		if len(password) >= 8 {
			break
		} else {
			fmt.Println("Incorrect password")
		}
	}

	if userType == "1" {
		return userType, firstName, lastName, phone, email, region, id, orphanageID, password
	}

	for {
		fmt.Printf("Passport ID (12 char): ")
		scanner.Scan()
		id = scanner.Text()
		if len(id) == 12 {
			break
		} else {
			fmt.Println("Incorrect ID")
		}
	}

	for {
		fmt.Printf("Orphanage name: ")
		scanner.Scan()
		orphanageID = scanner.Text()
		if len(orphanageID) == 0 {
			break
		} else {
			fmt.Println("Orphanage with ", id, " name does not exist")
		}
	}

	return userType, firstName, lastName, phone, email, region, id, orphanageID, password
}
