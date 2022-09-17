package main

import (
	"fmt"
	"unicode"
)

func main() {
	var testPassword string
	var testPasswordRune []rune
	var checkCorrectPassword = true
	fmt.Scan(&testPassword)

	for _, elem := range testPassword {
		testPasswordRune = append(testPasswordRune, elem)
	}

	if len(testPasswordRune) >= 5 {
		for _, elem := range testPasswordRune {
			if unicode.Is(unicode.Latin, elem) == false && unicode.Is(unicode.Digit, elem) == false {
				checkCorrectPassword = false
				break
			}
		}
	} else {
		checkCorrectPassword = false
	}

	if checkCorrectPassword == true {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
