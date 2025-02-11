package main

import (
	"fmt"
	"net/http"
)

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user! Welcome to our TODO List App!"
	fmt.Fprintf(writer, "%s", greeting)
}

func test(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user! This is /test!"
	fmt.Fprintf(writer, "%s", greeting)
}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/test", test)

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
