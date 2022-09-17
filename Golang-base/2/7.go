package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	var workRune []rune
	var checkPalindrom = true
	fmt.Scan(&text)

	text = strings.Trim(text, "\t\r\n ")

	for _, elem := range text {
		workRune = append(workRune, elem)
	}

	for i := 0; i < len(workRune); i++ {
		if workRune[i] != workRune[len(workRune)-1-i] {
			checkPalindrom = false
		}
	}
	if checkPalindrom == true {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
