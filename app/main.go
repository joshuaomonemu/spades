package main

import (
	"app/models"
)

func main() {
	haq := &models.Users{
		Firstname: "vokes",
	}
	models.GetUser(haq)
}
