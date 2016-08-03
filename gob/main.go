package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Test struct {
	Label string
	Type  int32
	Reps  []int64
	Man   *Person
}

type Person struct {
	Name string
}

func main() {
	test := &Test{
		Label: "Hello",
		Type:  17,
		Reps:  []int64{1, 2, 3},
		Man: &Person{
			Name: "Oliver",
		},
	}

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	dec := gob.NewDecoder(buf)

	err := enc.Encode(test)
	if err != nil {
		log.Fatalf("encode error: %v", err)
	}

	fmt.Printf("buf: %d\n%v\n\n", buf.Len(), buf.Bytes())

	newTest := new(Test)
	err = dec.Decode(newTest)
	if err != nil {
		log.Fatalf("decode error: %v", err)
	}

	fmt.Printf("test   : %s %d %v %s\n", test.Label, test.Type, test.Reps, test.Man.Name)
	fmt.Printf("newTest: %s %d %v %s\n", newTest.Label, newTest.Type, newTest.Reps, newTest.Man.Name)
}
