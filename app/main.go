package main

import (
	"app/models"
	"app/routes"
	"fmt"
)

func main() {
	//Route points
	vamp := models.Readmsgcontent()
	fmt.Println(vamp)
	routes.Routes()
}
