package models

import (
	"app/database"
	"log"

	firestore "cloud.google.com/go/firestore"
)

var msgClient = database.CreateClient().Collection("messages")
var details []string

//Messages payload Format
type Messages struct {
	Id        string `json:"messageId"`
	Content   string `json:"content"`
	Type      string `json:"type"`
	UserId    string `json:"userId"`
	Time      string `json:"time"`
	Recipient string `json:"recieverId"`
}

func CreateMsg(msg *Messages) bool {
	_, err := msgClient.Doc(msg.Id).Create(ctx, map[string]interface{}{
		"id":        msg.Id,
		"content":   msg.Content,
		"type":      msg.Type,
		"time":      msg.Time,
		"userId":    msg.UserId,
		"recipient": msg.Recipient,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return false
	}

	content := msg.Id
	vall := GetMessage(msg.UserId)
	val2 := GetMessage(msg.Recipient)
	vall = append(vall, content)
	val2 = append(val2, content)
	_, err = client.Doc(msg.Recipient).Set(ctx, map[string]interface{}{
		"messages": vall,
	}, firestore.MergeAll)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	_, err = client.Doc(msg.UserId).Set(ctx, map[string]interface{}{
		"messages": val2,
	}, firestore.MergeAll)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	return true

}

func DeleteMsg(key string) bool {
	_, err := msgClient.Doc("msg-01").Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}

func UpdateMsg() {

}

func ReadMsg(user1, user2 string) []string {
	vall := GetMessage(user1)
	val2 := GetMessage(user2)
	val := append(val2, vall...)
	for i, element := range val {
		for j, element2 := range val {
			if element == element2 && j > i {
				details = append(details, element)
			}
			for _, message := range details {

			}
		}
	}
	return details
}

// func MsgReader() {
// 	for _, v := range details {

// 	}
// }
