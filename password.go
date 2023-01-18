package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var strPas string
	strPas, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	strPasRunes := []rune(strPas)
	lenght := len(strPasRunes)
	var isValid bool
    for _, v := range strPasRunes {
        if !(unicode.Is(unicode.Latin, v) || unicode.Is(unicode.Digit, v)) || lenght < 5  {
				isValid = false
            break
			} else {
				isValid = true
			}
    }
	if isValid == true {
		fmt.Println("Ok")
	} else if isValid == false {
		fmt.Println("Wrong password")
	}

}
