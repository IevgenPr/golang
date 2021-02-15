package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers sets up app routing - receives all requests
func RegisterControllers() {
	uc := newUserController()
	// Handle(pattern string, ^^handler http.Handler) userController implements
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

}

func encodeResponseJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
