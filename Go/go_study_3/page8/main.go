package main

import "fmt"

func main() {
	totalScore := ask(1, "anjing")
	// Tambahkan nilai return dari function ask ke totalScore
	totalScore += ask(2, "kucing")
	totalScore += ask(3, "ikan")

	
	fmt.Println("Skor", totalScore)
	
}

func ask(number int, question string) int {
	var input string
	fmt.Printf("[Pertanyaan %d] Silahkan masukkan kata berikut: %s\n", number, question)
	fmt.Scan(&input)
	fmt.Printf("%s telah dimasukkan\n", input)
    
    return 10
	
}
