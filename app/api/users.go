package api

import (
	"app/models"
	"log"
	"net/http"
)

var req *http.Request

var sample = &models.Users{
	Id:       "galninvfi",
	Email:    req.FormValue("email"),
	Username: req.FormValue("username"),
	Password: req.FormValue("password"),
}

//Getting a particular user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {

	ss := string(models.ReadUser(sample.Id))
	if ss == "" {
		log.Fatalln("Sorry an error occured")
	}
	w.Header().Set("payload", ss)
}

//Creating new users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/index", http.StatusBadRequest)
	}
	response := models.CreateUser(sample)
	if response {
		w.Header().Set("success", "true")
	}
}

//Updating user details
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	key := "Falcone"
	response := models.UpdateUser(key)
	if response {
		w.Header().Set("Content-type", "application/json")
	}
}

//Deleting user details
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	key := "Falcone"
	response := models.DeleteUser(key)
	if response {
		w.Header().Set("Content-type", "application/json")
	}
}
