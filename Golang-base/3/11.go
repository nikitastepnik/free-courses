package main

import (
	"fmt"
	"time"
)

func main() {
	var sec, min string
	const now = 1589570165

	fmt.Scanf("%v мин. %v", &min, &sec)

	dur, err := time.ParseDuration(min + "m" + sec + "s")

	if err != nil {
		panic(err)
	}

	nowTime := time.Unix(now, 0)

	fmt.Println(nowTime.Add(dur).Format(time.UnixDate))

}
