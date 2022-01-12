package main

import "fmt"

func main() {
	// Tetapkan nilai return pada `totalScore`
	totalScore := ask(1, "anjing")

	// Tempelkan code yang diberikan dalam instruksi
	fmt.Println("Skor", totalScore)
	
}

// Atur tipe data dari nilai return
func ask(number int, question string) int {
	var input string
	fmt.Printf("[Pertanyaan %d] Silahkan memasukkan kata berikut: %s\n", number, question)
	fmt.Scan(&input)
	fmt.Printf("%s telah dimasukkan\n", input)
    // Hasilkan 10 sebagai nilai return
	return 10
	
}
