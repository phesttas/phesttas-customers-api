package main

import (
	"fmt"
	"net/http"
)

func createUserEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "createUser")
}
