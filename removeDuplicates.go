package main 

import (
	"fmt"
)

func removeDuplicates(inputStream, outputStream chan string) {
	var previous string 
	for v := range inputStream {
		if previous != v {
			outputStream <- v
			previous = v
		}
	}
	close(outputStream)
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	var input string
	fmt.Printf("Enter some text with duplicates: ")
	fmt.Scanln(&input)

	go func() {
		defer close(inputStream)

		for _, v := range input {
			inputStream <- string(v)
		}
	}()

	fmt.Printf("Output without duplicates: ")
	for v := range outputStream {
		fmt.Printf("%s", v)
	}
	fmt.Printf("\n")
}