package main

import (
	"fmt"
	"unsafe"
)

// MAX_SIZE is an untyped numeric constant
const MAX_SIZE = 1024

// VERSION is a typed string constant
const VERSION string = "1.0.0"

// use iota to mimic enum
const (
	TYPE_1 = 2 + iota
	TYPE_2
	TYPE_3
)

// Person is a structure
type Person struct {
	name string
	age  int
}

func main() {
	// -----------------------------------
	// Numeric Type
	// -----------------------------------
	// uint8, uint16, uint32, uint64
	// int8, int16, int32, int64
	// float32, float64
	// comlex64, complex128
	// byte = uint8
	// rune = int32

	// By default, a will be initialized with zero value
	var a int

	// You can also assign an initial value
	var b int = 2

	// Let the compiler decide the type (int in this case)
	var c = 3

	// Similar to var d = 4
	d := 4

	// Declare multiple variables
	var e, f int

	// Use () to declare multiple variables
	var (
		g     = 1
		h int = 2
	)

	// byte is an one-byte character
	var ch byte = 'A'

	// rune is a UTF-8 character
	var ch2 rune = '好'

	// Swap b and c
	b, c = c, b

	// Swap c and d
	swap(&c, &d)

	fmt.Println("--------- Numeric Type ----------")
	fmt.Println(a + b + c + d + e + f + g + h)
	fmt.Println(string(ch), string(ch2))
	fmt.Println("Type:", TYPE_1, TYPE_2, TYPE_3)

	// -----------------------------------
	// Boolean Type
	// -----------------------------------
	var isRight bool = true

	fmt.Println("--------- Boolean Type ----------")
	fmt.Println(isRight)

	// -----------------------------------
	// String Type
	// -----------------------------------
	var str1 string = "Hello"
	const str2 string = "World"
	str3 := "Hi你好"

	// Length is the number of bytes
	lenStr3 := len(str3)

	fmt.Println("--------- String Type ----------")
	fmt.Println(str1+str2, str3)
	fmt.Printf("%s: %d\n", str3, lenStr3)

	// Range string will get runes by runes (not bytes)
	for i, c := range str3 {
		fmt.Println(i, string(c))
	}

	// -----------------------------------
	// Array Type
	// -----------------------------------
	var array1 [2]int
	array2 := [...][3]int{{1, 2, 3}, {5, 6, 7}}
	array3 := [10]int{1: 22, 5: 11}
	var array4 [2]int = array1 // copy array
	array4[1] = 5
	var array5 [MAX_SIZE]byte

	// Call by value for array type
	func(array [2]int) { array[0] = 5 }(array1)

	fmt.Println("--------- Array Type ----------")
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(len(array3), array3)
	fmt.Println(array4)
	fmt.Println(len(array5), unsafe.Sizeof(array5))

	// -----------------------------------
	// Slice Type
	// -----------------------------------
	var slice1 []int               // nil slice
	slice2 := []int{}              // empty slice
	slice3 := make([]int, 3)       // len = cap = 3
	slice4 := make([]string, 3, 5) // len = 3, cap = 5
	slice5 := []string{"AA", "BBB", "CC", "EE", "FF"}
	slice6 := [][]int{}
	slice6 = append(slice6, []int{1, 2})
	slice6 = append(slice6, []int{4, 5, 6})
	slice7 := slice5[1:3]                 // {"BBB", "CC"}
	slice7 = append(slice7, "GG")         // slice5[3] = "GG"
	slice7 = append(slice7, "DDD", "EEE") // realloc, slice5 is not changed
	slice8 := make([]string, len(slice7))
	copy(slice8, slice7)
	slice9 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		slice9[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			slice9[i][j] = i + j
		}
	}

	fmt.Println("--------- Slice Type ----------")
	fmt.Println(unsafe.Sizeof(slice1), unsafe.Sizeof(slice2), slice1 == nil, slice2 == nil)
	fmt.Println(slice3)
	fmt.Println(slice4, len(slice4), cap(slice4))
	fmt.Println(slice5)
	fmt.Println(slice6)
	for i, v := range slice5 {
		fmt.Println(i, v)
	}
	fmt.Println(slice7)
	fmt.Println(slice8)
	fmt.Println(slice9)

	// -----------------------------------
	// Map Type
	// -----------------------------------
	dict := map[string]string{}
	//dict := make(map[string]string)
	dict["one"] = "single"
	dict["apple"] = "fruit"
	dict["oliver"] = "green arrow"
	delete(dict, "oliver")
	dict2 := map[string]int{"foo": 1, "bar": 2}

	fmt.Println("--------- Map Type ----------")
	fmt.Println("dict:", dict)
	fmt.Println("dict2:", dict2)
	value, exist := dict["one"]
	if exist {
		fmt.Println("Find", value)
	}

	_, exist = dict["two"]
	if !exist {
		fmt.Println("Couldn't find two")
	}

	for key, value := range dict {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// -----------------------------------
	// Channel Type
	// -----------------------------------
	end := make(chan int)
	go func() {
		end <- 5
	}()
	result := <-end
	close(end)

	fmt.Println("--------- Channel Type ----------")
	fmt.Println("Got channel end:", result)

	// -----------------------------------
	// Struct
	// -----------------------------------
	fmt.Println("--------- Struct ----------")
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 30})
	fmt.Println(Person{name: "Oliver"})
	var s Person
	// s := Person{"Bess", 25}
	s.name = "Bess"
	s.age = 25
	fmt.Println(s)

	ps := new(Person)
	// ps := &Person{}
	// var ps *Person = &Person{}
	ps.name = "John"
	ps.age = 11
	fmt.Println(ps)
}

// swap exchanges two integers
func swap(a, b *int) {
	*a, *b = *b, *a
}
