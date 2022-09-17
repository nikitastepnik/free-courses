package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type info struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type myStruct struct {
	ID       int
	Number   string
	Year     int
	Students []info
}

type averageStruct struct {
	Average float32
}

func main() {
	var data []byte
	var dataStructure myStruct
	var counts, sum int
	var average float32

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &dataStructure); err != nil {
		panic(err)
	}

	for _, elem := range dataStructure.Students {
		sum += len(elem.Rating)
		counts += 1
	}
	average = float32(sum) / float32(counts)

	averageStr := averageStruct{
		Average: average,
	}

	dataPrep, err := json.MarshalIndent(averageStr, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", dataPrep)
}
