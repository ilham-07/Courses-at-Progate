package main

import "fmt"

func main() {
	totalScore := 0
	fn(totalScore)

	// Cetak nilai dari variable `totalScore`
	fmt.Println(totalScore)
	// Cetak alamat dari variable `totalScore`
	fmt.Println(&totalScore)

}

func fn(totalScore int) {
	totalScore += 10
	// Cetak nilai dari variable `totalScore` pada function `fn`
	fmt.Println(totalScore)
	// Cetak alamat dari variable `totalScore` pada function `fn`
	fmt.Println(&totalScore)
}
