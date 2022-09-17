package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	text = strings.Trim(text, "\n")
	var textRune []rune
	for _, elem := range text {
		textRune = append(textRune, elem)
	}
	if unicode.IsUpper(textRune[0]) == true && textRune[len(textRune)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
