package api

import (
	"app/handlers"
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
func getUser(w http.ResponseWriter, r *http.Request) {

	ss := string(models.ReadUser(sample.Id))
	if ss == "" {
		log.Fatalln("Sorry an error occured")
	}
	w.Header().Set("payload", ss)
}

//Creating new users
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/index", http.StatusBadRequest)
	}
	response := models.CreateUser(sample)
	if response {
		w.Header().Set("success", "true")
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
	http.HandleFunc("/main", handlers.Home)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/api/get", getUser)
	http.HandleFunc("/api/create", createUser)
	http.HandleFunc("/api/update", updateUser)
	http.HandleFunc("/api/delete", deleteUser)
}
