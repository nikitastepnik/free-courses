package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var firstHalf, secondHalf string
	var firstNumb, secondNumb float64
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	firstHalf = strings.ReplaceAll(strings.Replace(strings.Split(input, ";")[0], ",", ".", 1),
		" ", "")
	secondHalf = strings.ReplaceAll(strings.Trim(strings.Replace(strings.Split(input, ";")[1], ",", ".",
		1), "\n"), " ", "")

	firstNumb, _ = strconv.ParseFloat(firstHalf, 64)
	secondNumb, _ = strconv.ParseFloat(secondHalf, 64)

	fmt.Printf("%.4f", firstNumb/secondNumb)
}
