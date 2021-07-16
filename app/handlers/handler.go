package handlers

import (
	"html/template"
	"log"
	"net/http"
)

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
func Active(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusMovedPermanently)
		return
	}
}

//handle logging out
func Logout(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "user_id",
		MaxAge: -1}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusFound)
}
