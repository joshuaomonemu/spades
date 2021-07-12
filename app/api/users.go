package api

import (
	"app/models"
	"fmt"
	"log"
	"net/http"
)

var sample = &models.Users{
	Id:        "80008h8hui",
	Firstname: "James",
	Lastname:  "Gordon",
	Username:  "jamy54",
	Birthdate: "24/4/1995",
}

//Getting a particular user by ID
func getUser(w http.ResponseWriter, res *http.Request) {
	ss := string(models.ReadUser(sample.Id))
	if ss == "" {
		log.Fatalln("Sorry an error occured")
	}
	fmt.Println(ss)

}

//Creating new users
func createUser(w http.ResponseWriter, r *http.Request) {
	respond := models.CreateUser(sample)
	if respond == "done" {
		w.Header().Set("response", respond)
	}
}

//Updating user details
func updateUser(w http.ResponseWriter, r *http.Request) {
	models.UpdateUser()
}

//Starting route points
func Routes() {
	http.HandleFunc("/api/users", getUser)
	http.HandleFunc("/api/create", createUser)
	http.HandleFunc("/api/update", updateUser)
}
