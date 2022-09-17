package main

func task2(chanel chan string, str string) {
	for i := 0; i < 5; i++ {
		chanel <- str + " "
	}
}
