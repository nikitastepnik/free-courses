package main

import "fmt"

func main() {

	var d int
	fmt.Scan(&d)

	degreePerHour := 360 / 12
	hours := d / degreePerHour
	minutes := d % degreePerHour * 60 / degreePerHour
	fmt.Println("It is", hours, "hours", minutes, "minutes.")
}
