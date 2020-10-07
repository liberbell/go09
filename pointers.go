package main

func fun(y int) {
	y = 0
}

func main() {
	y := 10
	func(y)
}