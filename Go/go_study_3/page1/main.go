package main

import "fmt"

func main() {
	totalScore := ask(1, "anjing")
	totalScore += ask(2, "kucing")
	totalScore += ask(3, "ikan")

	fmt.Println("Skor", totalScore)
}

func ask(number int, question string) int {
	var input string
	fmt.Printf("[Pertanyaan %d] Silahkan memasukkan kata berikut: %s\n", number, question)
	fmt.Scan(&input)

	score := 0

	if question == input {
		fmt.Println("Benar!")
		score = 10
	} else {
		fmt.Println("Salah!")
	}

	return score
}
