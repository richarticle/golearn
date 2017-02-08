package main

import (
	"fmt"
	"math"
	"reflect"
)

// Parent class
type shape interface {
	peri() float64
	area() float64
}

// Rectangle class
type rect struct {
	width, height float64
}

func (r rect) peri() float64 {
	return 2 * (r.width + r.height)
}

func (r rect) area() float64 {
	return r.width * r.height
}

// Circle class
type circle struct {
	radius float64
}

func (c circle) peri() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(s shape) {
	fmt.Println(reflect.TypeOf(s), s)
	fmt.Println("Perimeter:", s.peri())
	fmt.Println("Area:", s.area())
}

func main() {
	r := rect{2, 4}
	c := circle{3}

	measure(r)
	measure(c)

	s := []shape{rect{1, 2}, circle{1}}
	for _, v := range s {
		fmt.Printf("Area of %v%v = %f\n", reflect.TypeOf(v), v, v.area())
	}
}
