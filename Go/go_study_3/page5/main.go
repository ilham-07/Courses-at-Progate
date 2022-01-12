package main

import "fmt"

func main() {
	// Panggil function「ask」 sebanyak 3 kali dengan sebuah argument 
	ask("anjing")
	ask("kucing")
	ask("ikan")
	
}

// Lakukan edit pada function untuk menerima sebuah argument
func ask(question string) {
	var input string
	// Ubah code berikut ini dengan menggunakan `fmt.Printf`
	fmt.Printf("Silahkan memasukkan kata berikut: %s\n", question)
	fmt.Scan(&input)
	fmt.Printf("%s telah dimasukkan\n", input)
}