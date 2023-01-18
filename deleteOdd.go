package main

import (
	"fmt"
	"strconv"
)

func deleteOdd(i uint) uint {
		var s string
		a := strconv.FormatUint(uint64(i), 10)
		for _, v := range a {
			if (v-'0')%2 == 0 && (v-'0') != 0 {
				s += string(v)
			}
		}
		result, _ := strconv.ParseUint(s, 10, 64)
		if result == 0 {
			return uint(100)
		}
		return uint(result)
}

func main() {
	fmt.Println(deleteOdd(727178))

}