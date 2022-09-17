package main

import "fmt"

func main() {
	var seconds int
	fmt.Scan(&seconds)
	fmt.Printf("It is %d hours %d minutes.", seconds/3600, seconds%3600/60)
}
