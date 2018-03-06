package main

import "fmt"

func main() {
	var (
		x int
		y int
	)

	x = 4
	y = 9
	x, y = y, x
	fmt.Println((x), (y))
}
