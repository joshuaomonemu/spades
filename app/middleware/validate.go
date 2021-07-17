package middleware

import (
	"fmt"
	"math/rand"
	"net/http"
)

func ValidateEmail(w http.ResponseWriter, r *http.Request) string {
	email := r.FormValue("email")
	return email
}
func ValidateUsername(w http.ResponseWriter, r *http.Request) string {
	username := r.FormValue("username")
	return username
}
func ValidatePassword(w http.ResponseWriter, r *http.Request) string {
	password := r.FormValue("password")
	return password
}
func Identifier() string {
	ID := rand.Intn(1000)
	id := "spades" + fmt.Sprintf("%v", ID)
	return id

}
