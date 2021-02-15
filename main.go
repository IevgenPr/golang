package main

import (
	"net/http"

	"github.com/ipr0/golang/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil) // starts server

}
