package main

import "fmt"

func main() {
	var x, y float32
	var z int
	fmt.Scan(&x)
	z = int(x)
	y = float32(z) + 1
	for x != float32(y) {
		x = float32(int(x*10)+1) / 10
		fmt.Println(x)
	}
}
