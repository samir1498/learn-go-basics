package main

import (
	"fmt"
	"errors"
)

func add(a, b float64) float64 {
	return a + b
}

// Variadic function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
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
	r:= Rectangle{
		Width:  10,
		Height: 20,
	}
	fmt.Printf("Area: %.2f\n", r.Area())
	// Arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	// Arrays
	myInts := [3]int{1, 2, 3}
	fmt.Println(myInts)
	// Slices
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println(slice)
	// make slice with capacity make(slice, length, capacity)
	slice2 := make([]int, 3, 6)
	slice2[0] = 1
	slice2[1] = 2
	slice2[2] = 3
	fmt.Println(slice2)
	// slice literal
	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)
	// length of slice
	fmt.Println("length of slice:", len(mySlice))
	// capacity of slice
	fmt.Println("capacity of slice:", cap(slice2))
	// errors
	err := errors.New("error created")
	fmt.Println(err)
	// Variadic functions
	total := sum(1, 2, 3, 4, 5)
	fmt.Println(total)

}
