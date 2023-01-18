package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	date := scanner.Text()

	err := scanner.Err()
	if err != nil {
		panic(err)
	}

	t, err := time.Parse("2006-01-2 15:04:05", date)
	if err != nil {
		panic(err)
	}

	if t.Hour() < 13 {
		fmt.Println(t.Format("2006-01-2 15:04:05"))
	} else {
		t = t.Add(time.Hour * 24)
		fmt.Println(t.Format("2006-01-2 15:04:05"))
	}
}
