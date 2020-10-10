package main

import (
	"fmt"
)

func main() {
	x := 4
	fmt.Println("this is ", x)
	switch x {
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	}
}
