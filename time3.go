package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	err := scanner.Err()
	if err != nil {
		panic(err)
	}
	strSplited := strings.Split(str, ",")
	
	firstTime, err := time.Parse("02.01.2006 15:04:05", strSplited[0])
	if err != nil {
		panic(err)
	}
	secondTime, err := time.Parse("02.01.2006 15:04:05", strSplited[1])
	if err != nil {
		panic(err)
	}

	if firstTime.After(secondTime) {
		fmt.Println(firstTime.Sub(secondTime))
	} else {
		fmt.Println(secondTime.Sub(firstTime))
	}
}