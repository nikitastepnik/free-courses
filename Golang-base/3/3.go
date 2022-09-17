package main

import "strconv"

func main() {
	fn := func(i uint) uint {
		var finalNumb string
		for _, elem := range strconv.Itoa(int(i)) {
			elemAsInt, _ := strconv.ParseInt(string(elem), 10, 32)
			if elemAsInt%2 == 0 && elemAsInt != 0 {
				finalNumb += strconv.Itoa(int(elemAsInt))
			}
		}
		var finalInt, _ = strconv.Atoi(finalNumb)
		if finalInt == 0 {
			return 100
		} else {
			return uint(finalInt)
		}
	}

	fn(5)
}
