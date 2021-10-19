package routes

import (
	"app/api"
	"app/handlers"
	"app/util"
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
	r.HandleFunc("/test", handlers.Tester)

	// Api for User
	r.HandleFunc("/api/user/get", api.GetUser)
	r.HandleFunc("/api/user/create", api.CreateUser)
	r.HandleFunc("/api/user/update", api.UpdateUser)
	r.HandleFunc("/api/user/delete", api.DeleteUser)
	r.HandleFunc("/api/user/addfriend", api.Newfriend)

	// Api for messages
	r.HandleFunc("/api/message/create", api.CreateMessage)
	r.HandleFunc("/api/chat/get/", api.CreateMessage)

	r.HandleFunc("/ws", util.WsEndpoint)

	//Starting Server and running on port 2020
	http.ListenAndServe(":2020", nil)
}
