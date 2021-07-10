package models

import (
	"app/database"
	"context"
	"fmt"
	"log"
)

var client = database.CreateClient().Collection("users")
var ctx = context.Background()

type Users struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Birthdate string `json:"birthdate"`
}

//Creating a new User
func CreateUser() {
	_, err := client.Doc("userId").Create(ctx, map[string]interface{}{
		"firstname": "Los Angeles",
		"lastname":  "CA",
		"username":  "USA",
		"id":        "12345678",
		"birthdate": "nawa",
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
}

//Getting a single user by ID
func GetUser(userCall *Users) {
	data, err := client.Doc("alvinv#").Get(ctx)
	if err != nil {
		fmt.Println(err)
	}
	m := data.Data()

	//Storing Data Stored Respectively
	firstname := m["firstname"].(string)
	lastname := m["lastname"].(string)
	username := m["username"].(string)
	birthdate := m["birthdate"].(string)
	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(username)
	fmt.Println(birthdate)
}

//Updating a particular User by ID
func UpdateUser() {
	_, err := client.Doc("Falcone").Set(ctx, map[string]interface{}{
		"id":        "Dillidading",
		"firstname": "Carlus",
		"lastname":  "Capione",
		"username":  "qwerty54",
		"birthdate": "24/04/1997",
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

}

//Deleting User Document by ID
func DeleteUser() {
	_, err := client.Doc("Falcone").Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
}
