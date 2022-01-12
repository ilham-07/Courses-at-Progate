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
	fmt.Printf("[Pertanyaan %d] Silahkan masukkan kata berikut: %s\n", number, question)
	fmt.Scan(&input)

	if question == input {
	    // Tambahkan proses ketika kata dalam `question` sesuai dengan nilai `input`
		fmt.Println("Benar!")
		return 10
		
	} else {
	    // Tambahkan proses ketika kata dalam `question` tidak sesuai dengan nilai `input`
	    fmt.Println("Salah!")
		return 0
		
	}
}
