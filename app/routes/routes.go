package routes

import (
	"app/api"
	"app/handlers"
	"net/http"
)

//Starting route points
func Routes() {
	fs := http.FileServer(http.Dir("./view/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/main", handlers.Home)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/api/get", api.GetUser)
	http.HandleFunc("/api/create", api.CreateUser)
	http.HandleFunc("/api/update", api.UpdateUser)
	http.HandleFunc("/api/delete", api.DeleteUser)
}
