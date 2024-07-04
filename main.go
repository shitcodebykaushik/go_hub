package main

import (
    "fmt"
)

func main() {
    // Declare and initialize variables
    var name string = "World"
    age := 30
    pi := 3.14

    // Print a greeting message
    fmt.Println("Hello, ", name)

    // Print age
    fmt.Println("Age: ", age)

    // Print the value of pi
    fmt.Printf("Value of Pi: %.2f\n", pi)

    // Call a custom function
    result := add(5, 3)
    fmt.Println("5 + 3 =", result)
}

// Function to add two integers
func add(a int, b int) int {
    return a + b
}
