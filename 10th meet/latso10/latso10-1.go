package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	for x > 0 {
		x = x / 10
		y++
	}
	fmt.Print(y)
}
