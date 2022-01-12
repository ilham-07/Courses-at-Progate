package main

import "fmt"

func main() {
    totalScore := 0
	ask(1, "anjing", &totalScore)
	ask(2, "kucing", &totalScore)
	ask(3, "ikan", &totalScore)

	fmt.Println("Skor", totalScore)
}

func ask(number int, question string, scorePtr *int) {
	var input string
	fmt.Printf("[Pertanyaan %d] Silahkan memasukkan kata berikut: %s\n", number, question)
	fmt.Scan(&input)
	
    if question == input {
		fmt.Println("Benar!")
		*scorePtr += 10
	} else {
		fmt.Println("Salah!")
	}
}