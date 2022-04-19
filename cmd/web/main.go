package main

import (
	"fmt"
	"github.com/jjang65/go-hello-word/pkg/handlers"
	"net/http"
)

const portNumber = ":8081"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
