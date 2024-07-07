package main

import (
	"fmt"
	u "structs/user"
)

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBirthdate := getUserData("Enter your birthdate (MM/DD/YYYY): ")

	appUser, err := u.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println("An error occurred:\n", err)
		return
	}
	fmt.Println(appUser.GetUserDetails())

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
