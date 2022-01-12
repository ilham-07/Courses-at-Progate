package main

import "fmt"

func main() {
    // Deklarasikan variable input dengan tipe string
    var input string
    
    // Cetak "Silahkan memasukkan kata berikut: kucing"
    fmt.Println("Silahkan memasukkan kata berikut: kucing")
    
    // Gunakan Scan untuk membaca nilai input
    fmt.Scan(&input)
    
    // Cetak nilai input
    fmt.Printf("%s telah dimasukkan", input)
}