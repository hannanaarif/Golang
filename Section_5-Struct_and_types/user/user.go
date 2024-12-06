package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	Birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func (u User) OutPutDetails() {
	fmt.Println("firstName: ",u.firstName, "lastName",u.lastName, "Dateof Birth: ",u.Birthdate)
}

func (u *User) ClearUserName() {
	fmt.Println("clearing users name")
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email,password string) Admin{
	 return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "xyz",
			lastName: "okok",
			Birthdate: "---",
			createdAt: time.Now(),
		},
	 }
}

func New(firstName, lastName, Birthdate string) (*User,error) {
    if firstName==""||lastName==""||Birthdate==""{
		return nil,errors.New("First Name,last Name, and Birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		Birthdate: Birthdate,
		createdAt: time.Now(),
	},nil
}