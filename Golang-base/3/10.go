package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	buf, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	buf = strings.Trim(buf, "\n")

	str1, str2 := strings.Split(buf, ",")[0], strings.Split(buf, ",")[1]

	time1, err := time.Parse("02.01.2006 15:04:05", str1)
	if err != nil {
		panic(err)
	}

	time2, err := time.Parse("02.01.2006 15:04:05", str2)
	if err != nil {
		panic(err)
	}

	if time1.Before(time2) {
		fmt.Println(time2.Sub(time1))
	} else {
		fmt.Println(time1.Sub(time2))
	}
}
