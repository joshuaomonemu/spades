package api

import (
	"app/models"
	"encoding/json"
	"net/http"
	"net/url"
)

// Create Message
func CreateChat(w http.ResponseWriter, r *http.Request) {
	// FormData
	data := &models.Chats{
		Id:       r.FormValue("id"),
		Type:     r.FormValue("type"),
		Messages: r.FormValue("messages"),
		Members:  r.FormValue("members"),
		Date:     r.FormValue("date"),
	}

	response := models.CreateChat(data)

	if response {
		w.Header().Set("status", "200")
		w.Header().Set("response", "success")
	}
}

//Get Messages
func GetChat(w http.ResponseWriter, r *http.Request) {
	chatId := url.Values{}.Get("chatId")

	data, err := models.GetChat(chatId)

	if err != nil {
		w.Header().Set("status", "404")
		w.Header().Set("response", "success")
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}
