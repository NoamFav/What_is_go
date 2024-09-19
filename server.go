package main

import (
	"fmt"
	"net/http"
)

// helloHandler responds to requests with "Hello, World!"
func helloHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", helloHandler) // All requests are handled by helloHandler

	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil) // Start the server on port 8080
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
