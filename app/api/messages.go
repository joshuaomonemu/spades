package api

import (
	"app/models"
	"net/http"
)

// Create Message
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	// FormData
	data := &models.Messages{
		Content:   r.FormValue("content"),
		Type:      r.FormValue("type"),
		UserId:    r.FormValue("userId"),
		Recipient: r.FormValue("recipientId"),
		Time:      r.FormValue("time"),
	}

	response := models.CreateMsg(data)

	if response {
		w.Header().Set("status", "200")
		w.Header().Set("response", "success")
	}
}

//Get Messages
func GetMessage() {}
