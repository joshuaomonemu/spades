package main

import (
	"app/models"
	"app/routes"
	"fmt"
)

func main() {
	//Route points

	vamp := models.ReadMsg("life_21", "marypoppins")
	fmt.Println(vamp)
	routes.Routes()
}
