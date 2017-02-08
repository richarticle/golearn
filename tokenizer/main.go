package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str := "aa,  bb, ccc"

	// Split may get empty string (ex. double space)
	fmt.Println("strings.Split")
	tokens := strings.Split(str, " ")
	for i := range tokens {
		fmt.Println(i, tokens[i])
	}

	// Fields will not get empty string
	fmt.Println("\nstrings.Fields")
	tokens = strings.Fields(str)
	for i := range tokens {
		fmt.Println(i, tokens[i])
	}

	// FieldsFunc with customized separator
	fmt.Println("\nstrings.FieldsFunc")
	f := func(c rune) bool {
		return c == ',' || c == ' '
	}
	tokens = strings.FieldsFunc(str, f)
	for i := range tokens {
		fmt.Println(i, tokens[i])
	}

	byteStr := []byte(str)

	// Split for bytes
	fmt.Println("\nbytes.Split")
	byteTokens := bytes.Split(byteStr, []byte(" "))
	for i := range byteTokens {
		fmt.Println(i, string(byteTokens[i]))
	}

	// Fields for bytes
	fmt.Println("\nbytes.Fields")
	byteTokens = bytes.Fields(byteStr)
	for i := range byteTokens {
		fmt.Println(i, string(byteTokens[i]))
	}

	// FieldsFunc for bytes
	fmt.Println("\nbytes.FieldsFunc")
	byteTokens = bytes.FieldsFunc(byteStr, f)
	for i := range byteTokens {
		fmt.Println(i, string(byteTokens[i]))
	}
}
