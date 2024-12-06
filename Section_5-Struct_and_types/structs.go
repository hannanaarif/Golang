package main

import (
	"Section_5/user"
	"fmt"
)

func main() {
	userFirstName := getUserData("Please Enter the first name: ")
	userLastName := getUserData("Please Enter the Last name: ")
	userBirthDate := getUserData("Please Enter the DOB (MM/DD/YYYY): ")

	// appUser := user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	Birthdate: userBirthDate,
	// 	createdAt: time.Now(),
	// }

	//constructor to make user
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	admin := user.NewAdmin("xyz@gmail.com", "123456789")

	admin.OutPutDetails()
	admin.ClearUserName()
	appUser.ClearUserName()

	fmt.Println(admin)

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf("Actual address: %p\n", &appUser)              // Prints raw memory address
	// outPutDetails(&appUser)
	// fmt.Println(appUser.firstName, appUser.Birthdate)
	//appUser.OutPutDetails()
	//appUser.ClearUserName()
	//appUser.OutPutDetails()

	//fmt.Print(userFirstName ,userLastName ,userBirthDate);
}

// func outPutDetails(u *user) {
// 	fmt.Printf("Address in outPutDetails: %p\n", u)
// 	fmt.Println(u.firstName, u.lastName, u.Birthdate)
// }

func getUserData(value string) string {
	var userDeatils string
	fmt.Print(value)
	fmt.Scanln(&userDeatils)
	return userDeatils
}
