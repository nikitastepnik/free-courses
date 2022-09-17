package main

import (
	"fmt"
)

func main() {
	var inputStr string

	fmt.Scan(&inputStr)

	for _, elem := range inputStr {
		fmt.Print((int(elem) - '0') * (int(elem) - '0'))
	}

}
