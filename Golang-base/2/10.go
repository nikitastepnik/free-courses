package main

import (
	"fmt"
	"strings"
)

func main() {
	var testString string

	fmt.Scan(&testString)

	for _, elem := range testString {
		if strings.Count(testString, string(elem)) == 1 {
			fmt.Print(string(elem))
		}
	}

}
