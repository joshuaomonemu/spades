package main

import (
	"app/api"
	"net/http"
)

func main() {

	//Route points
	api.Routes()

	//Starting Server and running on port 2020
	http.ListenAndServe(":2020", nil)

}
