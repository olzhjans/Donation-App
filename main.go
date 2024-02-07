package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	fmt.Println("Log in - 1, sign up - 2")
	fmt.Printf("Type: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if input == "1" { //LOG IN
		signIn()
	} else if input == "2" { //SIGN UP
		signUp()
	} else {
		fmt.Printf("Error...")
		return
	}
}
