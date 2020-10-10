package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("It's the weekday.")
	default:
		fmt.Println("It's the weekend.")
	}
}
