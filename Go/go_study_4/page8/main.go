package main

import "fmt"

func main() {
	totalScore := 0
	// Berikan alamat dari variable `totalScore` sebagai argument ke-3
	ask(1, "anjing", &totalScore)
	ask(2, "kucing", &totalScore)
	ask(3, "ikan", &totalScore)

	fmt.Println("Skor", totalScore)
}

// Tambahkan sebuah parameter pointer `scorePtr` dengan tipe *int sebagai parameter ke-3
func ask(number int, question string, scorePtr *int) {
	var input string
	fmt.Printf("[Pertanyaan %d] Silahkan masukkan kata berikut: %s\n", number, question)
	fmt.Scan(&input)

	if question == input {
		fmt.Println("Benar!")
		// Lakukan perhitungan menggunakan dereferensi pointer
		*scorePtr += 10
	} else {
		fmt.Println("Salah!")
	}
}
