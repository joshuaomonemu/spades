package api

import (
	"app/handlers"
	"app/models"
	"log"
	"net/http"
)

var sample = &models.Users{
	Id:       "renzov",
	Email:    "Jammy@yahoo.com",
	Username: "Foxy15",
	Password: "24/4/1995",
}

//Getting a particular user by ID
func getUser(w http.ResponseWriter, res *http.Request) {
	ss := string(models.ReadUser(sample.Id))
	if ss == "" {
		log.Fatalln("Sorry an error occured")
	}
	w.Header().Set("payload", ss)
}

//Creating new users
func createUser(w http.ResponseWriter, r *http.Request) {
	response := models.CreateUser(sample)
	if response {
		w.Header().Set("Content-type", "application/json")
	}
}

//Updating user details
func updateUser(w http.ResponseWriter, r *http.Request) {
	key := "Falcone"
	response := models.UpdateUser(key)
	if response {
		w.Header().Set("Content-type", "application/json")
	}
}

//Deleting user details
func deleteUser(w http.ResponseWriter, r *http.Request) {
	key := "Falcone"
	response := models.DeleteUser(key)
	if response {
		w.Header().Set("Content-type", "application/json")
	}
}

//Starting route points
func Routes() {
	fs := http.FileServer(http.Dir("./view/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/api/get", getUser)
	http.HandleFunc("/api/create", createUser)
	http.HandleFunc("/api/update", updateUser)
	http.HandleFunc("/api/delete", deleteUser)
}
