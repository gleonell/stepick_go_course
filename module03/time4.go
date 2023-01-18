package main

import (
	"fmt"
	"time"
)

const now = 1589570165

func main() {
	var min, sec string 
	fmt.Scanf("%s мин. %s сек.", &min, &sec)
	inputPeriod := min + "m" + sec + "s"
	
	dur, err := time.ParseDuration(inputPeriod)
	if err != nil {
		panic(err)
	}
	
	seconds := dur.Seconds() + now 
	
	t := time.Unix(int64(seconds), 0).UTC()

	fmt.Println(t.Format(time.UnixDate))
}