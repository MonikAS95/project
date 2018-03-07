package main

import "fmt"

func swap(x, y *int) {
	var r = *x
	*x = *y
	*y = r
}

func main() {
	x := 9
	y := 7
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
