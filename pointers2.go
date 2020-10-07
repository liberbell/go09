package main

import "fmt"

func num(yPTR *int) {
	*yPTR = 5
}

func main() {
	yPTR := new(int)
	num(yPTR)
	fmt.Println(*yPTR)
}
