package models

import (
	"app/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"google.golang.org/api/iterator"
)

var client = database.CreateClient().Collection("users")
var ctx = context.Background()

//Users payload Format
type Users struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type Friend struct {
	id       string `json:"id"`
	username string `json:"username"`
}
type Info struct {
	FriendList []string
}

//Creating a new User
func CreateUser(us *Users) bool {
	_, err := client.Doc(us.Username).Create(ctx, map[string]interface{}{
		"email":    us.Email,
		"username": us.Username,
		"id":       us.Username + fmt.Sprintf("%v", (rand.Intn(999))),
		"password": us.Password,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}

//Getting a single user by ID
func ReadUser(key, pass string) []byte {
	data, err := client.Doc(key).Get(ctx)
	if err != nil {
		val := []byte("error")
		return val
	}
	m := data.Data()

	//Storing Data Stored Respectively
	id := m["id"].(string)
	email := m["email"].(string)
	username := m["username"].(string)
	password := m["password"].(string)

	if key != username || pass != password {
		val := []byte("error")
		return val
	}

	payload := &Users{
		Id:       id,
		Email:    email,
		Username: username,
		Password: password,
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
		return false
	}
	return true
}

//Deleting User Document by ID
func DeleteUser(key string) bool {
	_, err := client.Doc("Falcone").Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}

func AddFriend(key, username string) {
	err, _ := client.Doc(key).Collection("friends").Doc(username).Create(ctx, map[string]interface{}{
		"id":       "usd-81",
		"username": "V_OKES",
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		// return false
	}
	// return true
}

func FriendList() []string {
	c := &Info{
		FriendList: []string{},
	}
	var lis string
	var usd []string
	iter := client.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		m := doc.Data()
		lis = m["username"].(string)

		c.FriendList = append(c.FriendList, lis)
		usd = c.FriendList
	}
	return usd
}
