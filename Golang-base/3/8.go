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

	firstTime, err := time.Parse(time.RFC3339, buf)

	if err != nil {
		panic(err)
	}

	fmt.Println(firstTime.Format(time.UnixDate))
}
