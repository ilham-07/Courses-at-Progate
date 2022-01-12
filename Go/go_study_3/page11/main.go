package main

import "fmt"

func main() {
	totalScore := 0
	ask(1, "anjing", totalScore)
	ask(2, "kucing", totalScore)
	ask(3, "ikan", totalScore)

	fmt.Println("Skor", totalScore)
}

// Tambahkan totalScore dengan tipe data `int` sebagai parameter ke-3
func ask(number int, question string, totalScore int) {
	var input string
	fmt.Printf("[Pertanyaan %d] Silahkan masukkan kata berikut: %s\n", number, question)
	fmt.Scan(&input)

	if question == input {
		fmt.Println("Benar!")
		totalScore += 10
	} else {
		fmt.Println("Salah!")
	}
}
