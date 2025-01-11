package main

import "fmt"

func main() {
	var x, y, z, i int
	fmt.Scan(&x)
	for z < x {
		i++
		fmt.Scan(&y)
		z = z + y
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", i, y, z)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.", z, i)
}
