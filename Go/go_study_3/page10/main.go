package main

import "fmt"

func main() {
    totalScore := 0
	ask(1, "anjing")
	ask(2, "kucing")
	ask(3, "ikan")

	fmt.Println("Skor", totalScore)
}

func ask(number int, question string) {
	var input string
	fmt.Printf("[Pertanyaan %d] Silahkan masukkan kata berikut: %s\n", number, question)
	fmt.Scan(&input)

	if question == input {
		fmt.Println("Benar!")
		// Tambahkan 10 ke totalScore
		totalScore +=10
		
	} else {
		fmt.Println("Salah!")
	}
}
