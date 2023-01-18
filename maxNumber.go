// // package main
// // import "fmt"

// // func main()  {
// // 	array := [5]int{}
// // 	var a int
// // 	for i:=0; i < 5; i++{
// // 		fmt.Scan(&a)
// // 		array[i] = a
// // 	}
// // 	slice := array[0:5]
// // 	for i := 0; i < len(array); i++ {
// // 		for j := 1; j < len(slice) - 1; j++ {
// // 			if slice[j-1] > array[j] {
// // 				array[i] = slice[i]
// // 			}
// // 		}
// // 	}
// // 	fmt.Println(array)

// // }

package main

import "fmt"

func main()  {
	array := [5]int{}
	var a int
	fmt.Println("Enter numbers")
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
	max := array[0]
	for _, values := range array {
		if values > max {
			max = values
		}
	}
	fmt.Println("Max number:",max)

}

