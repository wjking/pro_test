package main

import (
	"fmt"
	"time"
)

func main() {
	var counter int64
	for {
		time.Sleep(time.Millisecond * 1000)
		counter++
		fmt.Printf("new counter=%d\n", counter)
	}
}
