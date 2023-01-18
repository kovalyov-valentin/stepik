package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	batteryCharge string
}

func (b *Battery) String() string {
	zero := strings.Count(b.batteryCharge, "0")
	one := strings.Count(b.batteryCharge, "1")
	result := "[" + strings.Repeat(" ", zero) + strings.Repeat("X", one) + "]"
	return result
}

func main() {
	var str string
	fmt.Scan(&str)
	batteryForTest := &Battery{
		batteryCharge: str,
	}
	fmt.Println(batteryForTest)
}