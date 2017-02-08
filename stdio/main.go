package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// bufio
	// ReadString -> ReadBytes -> ReadSlice
	// ReadLine -> ReadSlice
	// ReadString and ReadBytes may re-allocate buffer if buffer size is not enough
	// ReadLine and ReadSlice will not re-allocate buffer

	r := bufio.NewReader(os.Stdin)

	// ReadString
	fmt.Print("Enter: ")
	str, _ := r.ReadString('\n') // str consists of \n
	//str = strings.TrimRight(str, "\n")
	fmt.Println("ReadString:", str)

	// ReadBytes
	fmt.Print("Enter: ")
	line, _ := r.ReadBytes('\n') // line consists of \n
	fmt.Println("ReadBytes:", string(line))

	// ReadLine
	fmt.Print("Enter: ")
	line, _, _ = r.ReadLine() // line does not consist of \n
	fmt.Println("ReadLine:", string(line))

	// ReadSlice
	fmt.Print("Enter: ")
	line, _ = r.ReadSlice('\n') // line consists of \n
	fmt.Println("ReadSlice:", string(line))

	// Scanln
	fmt.Print("Enter: ")
	fmt.Scanln(&str) // aa bb -> aa
	fmt.Println("Scanln:", str)

	// Scanf
	fmt.Print("Enter: ")
	fmt.Scanf("%s", &str) // aa bb -> aa
	fmt.Println("Scanf:", str)

	// Scan
	fmt.Print("Enter: ")
	fmt.Scan(&str) // aa bb -> aa
	fmt.Println("Scan:", str)

	// Scanner
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("Text:", scanner.Text())
	}
}
