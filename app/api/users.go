package api

import (
	"app/models"
	"net/http"
)

//Getting a particular user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {

	//Form Data
	sample := &models.Users{
		Email:    r.FormValue("email"),
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	ss := string(models.ReadUser(sample.Username, sample.Password))
	if ss == "error" {
		w.Header().Set("success", "false")
	} else {
		w.Header().Set("success", "true")
		//w.Header().Set("payload", ss)
	}
}

//Creating new users
func CreateUser(w http.ResponseWriter, r *http.Request) {

	//Form Data
	sample := &models.Users{
		Email:    r.FormValue("email"),
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

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
func Newfriend(w http.ResponseWriter, r *http.Request) {
	key := "usd-81"
	username := "nana"
	models.AddFriend(key, username)
}
