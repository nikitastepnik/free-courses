package main

import (
	"fmt" // пакет используется для проверки ответа, не удаляйте его
	"os"  // пакет используется для проверки ответа, не удаляйте его
	"strings"
)

func readTask() (interface{}, interface{}, interface{}) {
	return 1.0, 1.0, "+"
}

func main() {
	value1, value2, operation := readTask()

	if _, ok := value1.(float64); ok {
	} else {
		fmt.Printf("value=%v: %T", value1, value1)
		os.Exit(0)
	}
	if _, ok := value2.(float64); ok {
	} else {
		fmt.Printf("value=%v: %T", value2, value2)
		os.Exit(0)
	} //"+", "-", "*", "/"
	if _, ok := operation.(string); ok {
		oper := operation.(string)
		if len(oper) != 1 || strings.Contains("+-/*", oper) == false {
			fmt.Println("неизвестная операция")
			os.Exit(0)
		}
	} else {
		fmt.Println("неизвестная операция")
		os.Exit(0)
	}

	val1 := value1.(float64)
	val2 := value2.(float64)
	oper := operation.(string)

	switch oper {
	case "+":
		fmt.Printf("%.4f", val1+val2)
	case "-":
		fmt.Printf("%.4f", val1-val2)
	case "*":
		fmt.Printf("%.4f", val1*val2)
	case "/":
		fmt.Printf("%.4f", val1/val2)
	}
}
