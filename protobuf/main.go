package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/richarticle/golearn/protobuf/protocol"
)

func main() {
	test := &protocol.Test{
		Label: "Hello",
		Type:  17,
		Reps:  []int64{1, 2, 3},
		Man: &protocol.Person{
			Name: "Oliver",
		},
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Printf("protobuf len = %d\n%v\n\n", len(data), data)
	newTest := &protocol.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if reflect.DeepEqual(test, newTest) == false {
		log.Fatal("data mismatch")
	}
	fmt.Println("newTest:", newTest.Label, newTest.Type, newTest.Reps, newTest.Man.Name)

	data, err = json.Marshal(test)
	fmt.Printf("\njson len = %d\n%s\n", len(data), data)
}
