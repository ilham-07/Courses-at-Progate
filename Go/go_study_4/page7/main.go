package main

import "fmt"

func main() {
	x := 10
	y := 10

	// Panggil function `calculate`, dan berikan x dan alamat dari variable y sebagai argument
	calculate(x, &y)

	fmt.Println("Ketika mengatur sebuah integer sebagai sebuah argument:", x)
	fmt.Println("Ketika mengatur sebuah pointer sebagai sebuah argument:", y)
}

func calculate(x int, yPtr *int) {
	// Deskripsikan proses menambahkan 1 pada variable x dan yPtr
	x += 1
	*yPtr += 1
}
