# Readme: Learning Go Basics

This is a simple Go program that covers fundamental concepts of the Go programming language. The code snippets provided in this package serve as a starting point for learning Go basics. The code includes examples of variables, functions, structs, interfaces, arrays, slices, error handling, and more.

## Getting Started

To run the code, make sure you have Go installed on your machine. You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).

Clone this repository and navigate to the directory containing the `main.go` file. Run the following command to execute the program:

```bash
go run main.go
```

This should print "Hello, World!" and output the results of various Go features.

## Code Overview

### Variables

The code demonstrates different ways of declaring and initializing variables:

```go
var a int = 10
b := "Hello"
c := 10.5
```

### Functions

The `add` function performs addition, and the `sum` function is a variadic function that calculates the sum of integers:

```go
fmt.Println(add(10, 20))
total := sum(1, 2, 3, 4, 5)
```

### Structs

A `Person` struct is defined and initialized:

```go
p := Person{
    Name: "John",
    Age:  30,
}
fmt.Println(p)
```

### Interfaces

An interface `Shape` and a struct `Rectangle` implementing the interface are introduced:

```go
type Shape interface {	
    Area() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}
```

### Arrays and Slices

Examples of arrays and slices are provided:

```go
// Arrays
var arr [3]int
arr[0] = 1
arr[1] = 2
arr[2] = 3

// Slices
slice := []int{1, 2, 3}
slice = append(slice, 4)
```

### Error Handling

An error is created and printed using the `errors` package:

```go
err := errors.New("error created")
fmt.Println(err)
```

### Learning Resources

The code is based on the content of the video [Learning Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=un6ZyFkqFKo&t=11542s). Feel free to refer to the video for more in-depth explanations and hands-on learning.

## Conclusion

This Go program provides a foundation for understanding essential concepts. Experiment with the code, modify it, and explore additional features of the Go programming language to enhance your learning experience.