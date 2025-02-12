package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Name  string
	Lover string
}

func main() {
	// map[KeyType]ValueType
	data := map[string]string{
		"name":  "Yunona",
		"lover": "Peter",
	}

	// json.Marshal()
	jsonString, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(jsonString)
	}

	// json.Unmarshal()
	go_obj := Data{}
	err = json.Unmarshal(jsonString, &go_obj)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Name: %s, Lover: %s", go_obj.Name, go_obj.Lover)
	}
}
