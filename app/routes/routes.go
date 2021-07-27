package routes

import (
	"app/api"
	"app/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

//Starting route points
func Routes() {

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./view/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.Handle("/", r)
	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/main", handlers.Home)
	r.HandleFunc("/register", handlers.Register)
	r.HandleFunc("/logout", handlers.Logout)
	r.HandleFunc("/api/get", api.GetUser)
	r.HandleFunc("/api/create", api.CreateUser)
	r.HandleFunc("/api/update", api.UpdateUser)
	r.HandleFunc("/api/delete", api.DeleteUser)
	r.HandleFunc("/api/friend", api.Newfriend)

	//Starting Server and running on port 2020
	http.ListenAndServe(":2020", nil)
}
