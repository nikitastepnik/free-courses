package main

import (
	"fmt"
)

func main() {
	var inputStr, newStr string

	fmt.Scan(&inputStr)

	for _, elem := range inputStr {
		newStr += string(elem) + "*"

	}

	fmt.Println(newStr[0 : len(newStr)-1])

}
