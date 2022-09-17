package main

import (
	"fmt"
	"strings"
)

func main() {
	var A, numberToDel string

	fmt.Scan(&A)
	fmt.Scan(&numberToDel)
	fmt.Println(strings.Replace(string(A), string(numberToDel), "", -1))
}
