package main

import (
	"fmt"
	"strings"
)

func main() {
	var testString, subsString string

	fmt.Scan(&testString)
	fmt.Scan(&subsString)

	if strings.Contains(testString, subsString) == false || len(subsString) > len(testString) {
		fmt.Println(-1)
	} else {
		var lengthSubs = len(subsString)
		for i := 0; i < len(testString)+1; i++ {
			if testString[i:lengthSubs+i] == subsString {
				fmt.Println(i)
				break
			}
		}
	}

}
