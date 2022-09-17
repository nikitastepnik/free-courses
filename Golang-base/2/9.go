package main

import (
	"fmt"
)

func main() {
	var testString, newString string

	fmt.Scan(&testString)

	for idx, elem := range testString {
		if idx%2 == 1 {
			newString += string(elem)
		}
	}
	fmt.Println(newString)

}
