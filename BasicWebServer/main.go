package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TodoItem struct {
	Title       string
	Description string
	Done        bool
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user! Welcome to our TODO List App!"
	fmt.Fprintf(writer, "%s", greeting)
}

func test(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user! This is /test!"
	fmt.Fprintf(writer, "%s", greeting)
}

func todoItem(writer http.ResponseWriter, request *http.Request) {
	item := TodoItem{Title: "Watch Go crash course", Description: "Learning about Go lang", Done: false}
	fmt.Fprintf(writer, "Title: %s, Description: %s, Done: %t", item.Title, item.Description, item.Done)
}

func todoList(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read JSON body
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Failed to read request body", http.StatusBadRequest)
	}
	defer request.Body.Close()

	// Parse JSON
	todo := TodoItem{}
	err = json.Unmarshal(body, &todo)
	if err != nil {
		http.Error(writer, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Print received data
	fmt.Println("Received data:")
	fmt.Printf("Title: \"%s\", Description: \"%s\", Done: %t\n", todo.Title, todo.Description, todo.Done)

	// Respond with a success message
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "TODO item received successfully",
	}
	json.NewEncoder(writer).Encode(response)
}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/test", test)
	http.HandleFunc("/item", todoItem)
	http.HandleFunc("/list", todoList)

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
