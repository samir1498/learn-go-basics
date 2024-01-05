package main

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

// Define a struct
type Person struct {
    Name string
    Age  int
}

// Define an interface
type Shape interface {	
	Area() float64
}

// Define a struct that implements the interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Hello, World!
	fmt.Println("Hello, World!")
	// variables
	// verbose declaration
	var a int = 10
	// short declaration
	b := "Hello"
	c := 10.5
	// sprintf
	d := fmt.Sprintf("a: %d, b: %s, c: %.2f", a, b, c)
	// print
	fmt.Println(d)
	// functions
	fmt.Println(add(10, 20))
	// casting
	fmt.Println(add(float64(a), c))
	// structs
	p := Person{
		Name: "John",
		Age:  30,
	}
	fmt.Println(p)
	fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
	// Interfaces

}
