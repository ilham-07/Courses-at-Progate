package main

import "fmt"

func main() {
	ask()
	// Panggil function `ask` lagi
	ask()
}

func ask() {
    var input string
	fmt.Println("Silahkan memasukkan kata berikut: kucing")
    fmt.Scan(&input)
    fmt.Printf("%s telah dimasukkan\n", input)
}
