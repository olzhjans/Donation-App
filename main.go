package main

func main() {
	apiLaunch()

	/*var input string
	for input != "0" {
		fmt.Println("Log in - 1, sign up - 2, Show all orphanage's data - 3, API - 4")
		fmt.Printf("Type: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		switch input {
		case "1":
			userID, isAdmin := logIn()
			fmt.Println("Change password - 1, Edit orphanage data - 2")
			fmt.Printf("Type: ")
			scanner.Scan()
			input = scanner.Text()
			switch input {
			case "1":
				editUserData(userID, isAdmin)
			case "2":
				if isAdmin {
					fmt.Println("Enter orphanage name to edit its data")
					scanner.Scan()
					input = scanner.Text()
					editOrphanageData(input)
				}
			}
		case "2":
			signUp()
		case "3":
			fmt.Printf("%s\n", getOrphanageInfo(""))
		case "4":
			apiLaunch(getOrphanageInfo(""))
		default:
			fmt.Printf("Error...")
			return
		}
	}
	fmt.Println("Exit...")
	*/
}
