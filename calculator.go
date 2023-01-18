package main

import (
	"fmt"
)

func divide(a, b int) int {
	return a / b
}

func multiplication(a, b int) int {
	return a * b
}

func addition(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func main() {
	var first, second int
    var operator string
	fmt.Println("Enter first number")
	_, err := fmt.Scan(&first)
	if err != nil {
        fmt.Println("Incorrect type")
	}
    fmt.Println("Enter operator")
    _, err = fmt.Scan(&operator)
	fmt.Println("Enter second")
	_, err = fmt.Scan(&second) 
    if err != nil {
        fmt.Println("Incorrect type")
    }
    if operator == "/" {
        fmt.Println("Output:", divide(first, second))
    } else if operator == "*" {
        fmt.Println("Output:", multiplication(first, second))
    } else if operator == "+" {
        fmt.Println("Output:", addition(first, second))
    } else if operator == "-" {
        fmt.Println("Output:", subtraction(first, second))
    } else {
        fmt.Println("Incorrect operator")
    }
}
