package main 

import (
	"fmt"
	_"encoding/json"
	_"os"
)

func readTask() (value1, value2, operation interface{}) {
	return 5.0, 5.6, "/"
}

func printError(value interface{}) {
    fmt.Printf("value=%v: %T", value, value)

}

func main() {
	value1, value2, operation := readTask()

	v1, ok := value1.(float64)
	if !ok {
		printError(value1)
		return
	}

	v2, ok := value2.(float64)
	if !ok {
		printError(value2)
		return
	}

	v, ok := operation.(string)
	if ok {
		switch v {
		case "+":
			result := v1 + v2
			fmt.Printf("%.4f", result)
		case "-":
			result := v1 - v2
			fmt.Printf("%.4f", result)
		case "*":
			result := v1 * v2
			fmt.Printf("%.4f", result)
		case "/":
			result := v1 / v2
			fmt.Printf("%.4f", result)
		default:
			fmt.Println("неизвестная операция")
			return
		}
	} else {
		fmt.Println("неизвестная операция")
	}
}