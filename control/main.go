package main

import (
	"fmt"
	"time"
)

func main() {
	// -----------------------------------
	// If
	// -----------------------------------
	value := true
	if value {
		fmt.Println("value: true")
	} else {
		fmt.Println("value: false")
	}

	if num := 6; num%4 == 0 {
		fmt.Printf("%d %% 4 = 0\n", num)
	} else if num%3 == 0 {
		fmt.Printf("%d %% 3 = 0\n", num)
	} else {
		fmt.Printf("Here\n")
	}

	// -----------------------------------
	// For
	// -----------------------------------
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("")

	var i int32 = 0
	for i <= 5 {
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Println("")
	for {
		fmt.Println("Infinite loop")
		break
	}

	// -----------------------------------
	// Switch
	// -----------------------------------
	day := 3
	fmt.Printf("%d is ", day)
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Holycrap")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Before noon")
	default:
		fmt.Println("After noon")
	}
}
