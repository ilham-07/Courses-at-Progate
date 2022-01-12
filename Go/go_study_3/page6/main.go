package main

import "fmt"

func main() {
	// Berikan angka 1 sebagai sebuah argument baru
	ask(1, "anjing")
	
	// Berikan angka 2 sebagai sebuah argument baru
	ask(2, "kucing")
	
	// Berikan angka 3 sebagai sebuah argument baru
	ask(3, "ikan")
}

// Lakukan edit pada method untuk menerima `number` sebagai sebuah argument
func ask(number int, question string) {
	var input string
	// Hapus 1 baris di bawah ini dan tempelkan code yang diberikan
	fmt.Printf("[Pertanyaan %d] Silahkan memasukkan kata berikut: %s\n", number, question)
	
	fmt.Scan(&input)
	fmt.Printf("%s telah dimasukkan\n", input)
}
