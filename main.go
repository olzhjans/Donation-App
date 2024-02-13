package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	for input != "0" {
		fmt.Println("Log in - 1, sign up - 2")
		fmt.Printf("Type: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()

		switch input {
		case "1":
			if isModerator := signIn(); isModerator {
				fmt.Println("Enter orphanage name to edit its data")
				scanner.Scan()
				input = scanner.Text()
				editOrphanageData(input)
			}
		case "2":
			signUp()
		case "3":
			getOrphanageInfo("")
		default:
			fmt.Printf("Error...")
			return
		}
	}
	fmt.Println("Exit...")
}
