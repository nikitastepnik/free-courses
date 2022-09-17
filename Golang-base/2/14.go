package main

import (
	"fmt"
)

func main() {
	var inputStr, maxNumber string

	fmt.Scan(&inputStr)
	maxNumber = string(rune(inputStr[0]))
	for i := 1; i < len(inputStr); i++ {
		if maxNumber < string(rune(inputStr[i])) {
			maxNumber = string(rune(inputStr[i]))
		}
	}

	fmt.Println(maxNumber)

}
