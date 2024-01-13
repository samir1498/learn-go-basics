package main

import (
	"errors"
	"fmt"
)

// CustomError is a custom error type for demonstration purposes
type CustomError struct {
	message string
}

func (e CustomError) Error() string {
	return e.message
}

func add(a, b float64) (float64, error) {
	return a + b, nil
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

	// Variables
	a := 10
	b := "Hello"
	c := 10.5

	// Printf
	d := fmt.Sprintf("a: %d, b: %s, c: %.2f", a, b, c)
	fmt.Println(d)

	// Functions
	result, err := add(10, 20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum:", result)
	}

	// Casting
	result, err = add(float64(a), c)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum:", result)
	}

	// Structs
	p := Person{
		Name: "John",
		Age:  30,
	}
	fmt.Println(p)
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)

	// Interfaces
	r := Rectangle{
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

	// Make slice with capacity make(slice, length, capacity)
	slice2 := make([]int, 3, 6)
	slice2[0] = 1
	slice2[1] = 2
	slice2[2] = 3
	fmt.Println(slice2)

	// Slice literal
	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)

	// Length of slice
	fmt.Println("Length of slice:", len(mySlice))

	// Capacity of slice
	fmt.Println("Capacity of slice:", cap(slice2))

	// Errors
	err = errors.New("error created")
	fmt.Println(err)

	// Custom Error
	customErr := CustomError{"custom error message"}
	fmt.Println("Custom Error:", customErr)

	// Variadic functions
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers:", total)
}
