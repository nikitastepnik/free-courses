package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var sum int
	str := bufio.NewScanner(os.Stdin)

	for str.Scan() {
		if len(str.Text()) != 0 {
			numbInt, _ := strconv.Atoi(str.Text())
			sum += numbInt
		} else {
			break
		}
	}
	w := bufio.NewWriter(os.Stdout)
	_, err := w.WriteString(strconv.Itoa(sum))
	if err != nil {
		panic(err)
	}
	err1 := w.Flush()
	if err1 != nil {
		panic(err1)
	}
}
