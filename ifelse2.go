package main

import (
	"fmt"
	"time"
)

func main() {
	switch {
	case time.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("afternoon")
	}
}
