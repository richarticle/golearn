package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"-"`
	Age  int32  `json:"age,omitempty"`
}

func main() {
	var person Person

	// Unmarshal: []byte -> struct
	// Name will not be set
	input := []byte(`{"id":1,"name":"John","age":1}`)
	if err := json.Unmarshal(input, &person); err != nil {
		panic(err)
	}

	// Marshal: struct -> []byte
	// Age will not show
	person.Age = 0
	output, _ := json.Marshal(person)
	fmt.Println(string(output))
}
