package main

import (
	"fmt"
	"strings"
	"unicode"
	"strconv"
)


func adding(a, b string) int64 {

	aArr := []rune(a)
	bArr := []rune(b)

	for _, v := range aArr {
		if !unicode.IsDigit(v) {
			a = strings.ReplaceAll(a, string(v), "")
		}
	}
	first, _ := strconv.ParseInt(a, 10, 64)


	for _, v := range bArr {
		if !unicode.IsDigit(v) {
			b = strings.ReplaceAll(b, string(v), "")
		}
	}

	second, _ := strconv.ParseInt(b, 10, 64)
	return first + second

}
//%^80 hhhhh20&&&&nd
func main() {
	var a, b string 
	fmt.Scan(&a, &b)

	fmt.Println(adding(a, b))
}