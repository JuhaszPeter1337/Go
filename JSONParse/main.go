package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Lover string
	Age   int
}

func main() {
	// Raw JSON string
	data := `{"name": "Yunona", "lover": "Peter", "age": 24}`

	fmt.Println(data)

	// JSON parse
	person := Person{}
	err := json.Unmarshal([]byte(data), &person)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Name: %s, lover: %s, age: %d\n", person.Name, person.Lover, person.Age)
	}
}
