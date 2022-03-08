package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var ue UserEndpoint
	http.HandleFunc("/", ue.Find)

	fmt.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
