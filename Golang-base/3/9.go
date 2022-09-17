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

	now, err := time.Parse("2006-01-02 15:04:05", buf)

	if err != nil {
		panic(err)
	}

	if now.Hour() < 13 {
		fmt.Println(now.Format("2006-01-02 15:04:05"))
	} else {
		future := now.Add(time.Hour * 24)

		fmt.Println(future.Format("2006-01-02 15:04:05"))
	}

}
