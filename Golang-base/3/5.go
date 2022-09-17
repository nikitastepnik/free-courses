package main

import "fmt" // пакет используется для проверки ответа, не удаляйте его

type MyBattery struct {
	batteryString string
}

func (battery MyBattery) String() string {
	var resultStr = "["
	var countSpaces, countUnits int

	for _, elem := range battery.batteryString {
		if string(elem) == "1" {
			countUnits += 1
		} else {
			countSpaces += 1
		}
	}

	for i := 0; i < countSpaces; i++ {
		resultStr += " "
	}
	for i := 0; i < countUnits; i++ {
		resultStr += "X"
	}
	resultStr += "]"

	return resultStr
}

func main() {
	var inititialStr string

	fmt.Scan(&inititialStr)

	var batteryForTest MyBattery

	batteryForTest.batteryString = inititialStr
}
