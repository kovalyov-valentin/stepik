package main 

import (
	"fmt"
)

func sumInt(n ...int) (lenght int, sum int) {
    for _, value := range n {
        sum += value
        lenght++
    }
    return lenght, sum 
}

func main() {
	len, s := sumInt(1, 23, 15, 87, 7, 92)
	fmt.Println("Lenght:", len)
	fmt.Println("Sum:", s)
}