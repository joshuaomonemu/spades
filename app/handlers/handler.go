package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var req *http.Request

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("view/login.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, nil)
}
func Register(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("view/register.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, nil)
}

//Checking if sign in session is active
func Active() bool {
	_, err := req.Cookie("user_id")
	if err != nil {
		fmt.Println(err)
	}
	return true
}

//Handling the home page
func Home(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("view/main.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, nil)
}

//handle logging out
func Logout(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "user_id",
		MaxAge: -1}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusFound)
}
