package main

import (
	"app/routes"
	"net/http"
)

func main() {

	//Route points
	routes.Routes()

	//Starting Server and running on port 2020
	http.ListenAndServe(":2020", nil)

}
