package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("afternoon")
	}
}
