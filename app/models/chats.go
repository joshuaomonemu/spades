package models

import (
	"app/database"
	"encoding/json"
	"log"
)

var chatClient = database.CreateClient().Collection("messages")

type Chats struct {
	Id       string   `json:"id"`
	Date     string   `json:"date"`
	Type     string   `json:"type"`
	Members  []string `json:"members"`
	Messages []string `json:"messages"`
}

func CreateChat(chat *Chats) bool {
	_, err := chatClient.Doc(chat.Id).Create(ctx, map[string]interface{}{
		"id":       chat.Id,
		"members":  chat.Members,
		"type":     chat.Type,
		"messages": chat.Messages,
		"date":     chat.Date,
	})
	if err != nil {
		// Handle the errors
		log.Printf("An Error occurred: %s", err)
		return false
	}
	return true
}

func GetChat(key string) (interface{}, error) {
	// Return result and error status
	data, err := chatClient.Doc(key).Get(ctx)
	c := data.Data()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	payload := &Chats{
		Id:       c["id"].(string),
		Members:  c["members"].([]string),
		Type:     c["type"].(string),
		Messages: c["message"].([]string),
		Date:     c["date"].(string),
	}

	chat, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("Error converting to JSON")
	}
	return chat, nil
}
