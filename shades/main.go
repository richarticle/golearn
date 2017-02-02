package main

import (
	"errors"
	"fmt"
	"unsafe"
)

type ShadeFunc func()

var Shades map[int]ShadeFunc

func main() {

	RegisterShades()

	for {
		var n int
		fmt.Println("Which shade to see?")
		fmt.Scanf("%d", &n)
		if n <= 0 {
			fmt.Println("Bye")
			break
		}
		shade, ok := Shades[n]
		if !ok {
			fmt.Printf("Shade %d does not exist\n", n)
			continue
		}
		fmt.Printf("Shade %d\n", n)
		shade()

		fmt.Println()
	}
}

// RegisterShades collects all shade functions in a map
func RegisterShades() {
	Shades = map[int]ShadeFunc{}
	Shades[1] = Shade1
	Shades[2] = Shade2
	Shades[3] = Shade3
	Shades[4] = Shade4
	Shades[5] = Shade5
	Shades[6] = Shade6
	Shades[7] = Shade7
	Shades[8] = Shade8
	Shades[9] = Shade9
	Shades[10] = Shade10
	Shades[11] = Shade11
	Shades[12] = Shade12
	Shades[13] = Shade13
	Shades[14] = Shade14
}

// Shade 1: Remember to use "call by reference" for slice arguments
// if the slice may be re-alloacted.
func Shade1() {
	sl := []int{1, 2}

	badPush(sl, 3)
	fmt.Println("Bad push", sl)

	goodPush(&sl, 3)
	fmt.Println("Good push", sl)
}

func badPush(slice []int, n int) {
	slice = append(slice, n)
}

func goodPush(slice *[]int, n int) {
	*slice = append(*slice, n)
}

// Shade 2: Slice will be re-allocated if the capacity is not enough.
// The old space will keep the same if it is still referred by anyone.
func Shade2() {
	array := [3]int{1, 2}
	sl := array[0:2]
	sl = append(sl, 3)
	sl[0] = 9
	sl = append(sl, 4)
	sl[1] = 8
	fmt.Println("array became", array)
}

// Shade 3: Struct assignment copies the first-layer members. For slice member,
// only the reference is copied while the allocated spaces are the same.
func Shade3() {
	type T struct {
		a int
		b []int
	}

	A := T{1, []int{2, 3}}
	B := A
	B.a = 2
	B.b[0] = 4
	fmt.Printf("A: %v\nB: %v\n", A, B)
}

// Shade 4: Sub-slice's capacity depends on the original slice's capacity.
// When doing append, the space will not be re-allocated if the capacity
// is enough.
func Shade4() {
	sl := []int{1, 2, 3, 4}
	a := sl[0:2]
	b := sl[2:]
	a = append(a, 9)
	fmt.Println("b[0] became", b[0])
}

// Shade 5: defer's order is FILO
func Shade5() {
	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

// Shade 6: Variable shadowing may occur for return variables.
func Shade6() {
	result, _ := div(6, 3)
	fmt.Println("6/3 =", result)
}

func div(a, b int) (result int, hasError bool) {
	if result, err := divWithError(a, b); err != nil {
		fmt.Println(result)
		hasError = true
	}
	return
}

func divWithError(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divide By Zero")
	} else {
		return a / b, nil
	}
}

// Shade 7: Referring to for loop variable is dangerous.
func Shade7() {
	sl := []int{0, 1, 4, 9}
	square := []func() int{}

	for i := range sl {
		square = append(square, func() int {
			return sl[i]
		})
	}

	// Solution 1: Declare local variable inside loop
	/*
		for i := range sl {
			j := i
			square = append(square, func() int {
				return sl[j]
			})
		}
	*/

	// Solution 2: Use generating function
	/*
		for i := range sl {
			square = append(square, func(n int) func() int {
				return func() int {
					return sl[n]
				}
			}(i))
		}
	*/

	fmt.Println("1^2 =", square[1]())
	fmt.Println("2^2 =", square[2]())
	fmt.Println("3^2 =", square[3]())
}

// Shade 8: Structure alignment is important for memory efficiency
func Shade8() {
	type TypeA struct {
		a byte
		b byte
		c int32
	}
	type TypeB struct {
		a byte
		c int32
		b byte
	}
	var A TypeA
	var B TypeB
	fmt.Println("Sizeof A:", unsafe.Sizeof(A))
	fmt.Println("Sizeof B:", unsafe.Sizeof(B))
}

// Shade 9: String is constant and type conversion will make a copy.
func Shade9() {
	str := "hello"
	b := []byte(str)
	b[0] = 'G'
	fmt.Println(str)
}

// Shade 10: Deadlock will occur since writing 1 to ch will be blocked forever.
// You can make ch a buffered channel to fix the problem.
func Shade10() {
	//ch := make(chan int, 1)
	ch := make(chan int)
	ch <- 1
	result := <-ch
	fmt.Println(result)
}

// Shade 11: As an argument, array is called by value while slice is called
// by reference.
func Shade11() {
	x := [2]int{1, 2}
	func(arr [2]int) {
		arr[0] = 9
	}(x)
	fmt.Println("x[0] =", x[0])

	y := []int{1, 2}
	func(sl []int) {
		sl[0] = 9
	}(y)
	fmt.Println("y[0] =", y[0])
}

// Shade 12: Be cautious about variable reference in generating function
func Shade12() {
	n := 5
	f := getFunc(n)
	fmt.Println(f())
	fmt.Println(f())
	f = getFunc(n)
	fmt.Println(f())
}

func getFunc(n int) func() int {
	return func() int {
		n++
		return n
	}
}

// Shade 13: rect does not implement the shape interface but *rect does.
func Shade13() {
	//var s shape = rect{2, 4}
	var s shape = &rect{2, 4}
	fmt.Println(s.peri())
}

type shape interface {
	peri() float64
}

type rect struct {
	width  float64
	height float64
}

func (r *rect) peri() float64 {
	return 2 * (r.width + r.height)
}

// Shade 14: Variable reference for different types of defer
func Shade14() {
	i := 1
	defer func() {
		fmt.Println(i)
	}()
	defer fmt.Println(i)
	i = 2
}
