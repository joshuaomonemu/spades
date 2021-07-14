package models

import (
	"app/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

var client = database.CreateClient().Collection("users")
var ctx = context.Background()

//Users payload Format
type Users struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Birthdate string `json:"birthdate"`
}

//Creating a new User
func CreateUser(us *Users) bool {
	_, err := client.Doc(us.Id).Create(ctx, map[string]interface{}{
		"firstname": us.Firstname,
		"lastname":  us.Lastname,
		"username":  us.Username,
		"id":        us.Id,
		"birthdate": us.Birthdate,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	return true
}

//Getting a single user by ID
func ReadUser(key string) []byte {
	data, err := client.Doc("alvinv#").Get(ctx)
	if err != nil {
		fmt.Println(err)
	}
	m := data.Data()

	//Storing Data Stored Respectively
	id := m["id"].(string)
	firstname := m["firstname"].(string)
	lastname := m["lastname"].(string)
	username := m["username"].(string)
	birthdate := m["birthdate"].(string)

	payload := &Users{
		Id:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Username:  username,
		Birthdate: birthdate,
	}
	bs, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("Error converting to JSON")
	}
	return bs
}

//Updating a particular User by ID
func UpdateUser(key string) bool {
	_, err := client.Doc(key).Set(ctx, map[string]interface{}{
		"id":        "Dillidading",
		"firstname": "Carlus",
		"username":  "qwerty54",
		"birthdate": "24/04/1997",
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	return true
}

//Deleting User Document by ID
func DeleteUser(key string) bool {
	_, err := client.Doc("Falcone").Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	return true
}
