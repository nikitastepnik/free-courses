package main

import (
	"strconv"
	"unicode"
)

func adding(a, b string) int64 {
	strArray := [2]string{a, b}
	var helpVar string
	var workArray [2]int
	for i := 0; i < 2; i++ {
		_, err := strconv.Atoi(strArray[i])
		if err != nil {
			for _, elem := range strArray[i] {
				if unicode.Is(unicode.Digit, elem) == true {
					helpVar += string(elem)
				}
			}
			workArray[i], _ = strconv.Atoi(helpVar)
		} else {
			workArray[i], _ = strconv.Atoi(strArray[i])
		}
		helpVar = ""
	}

	return int64(workArray[0] + workArray[1])
}