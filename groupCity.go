package main 

import (
	"fmt"
)

func main() {
	groupCity := map[int][]string{
		10: []string{"ПГТ"},
		100: []string{"Город"}, 
		1000: []string{"Миллионник"},
	}

	cityPopulation := map[string]int {
		"Село":          10,
		"cело":          100,
		"Миллионик":     1000,
		"Город":         100,
		"Большой город": 100,
	}
	for key := range cityPopulation {
		var exist bool 
		for _, value := range groupCity[100] {
			if value == key {
				exist = true 
				break
			}
		}
		if !exist {
			delete(cityPopulation, key)
		}
	}
	fmt.Println(cityPopulation)

	// for key, cities := range groupCity {
	// 	for _, city := range cities {
	// 		if key == 100 {
	// 			continue
	// 		} else {
	// 			for k := range cityPopulation {
	// 				if k == city {
	// 					delete(cityPopulation, k)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}