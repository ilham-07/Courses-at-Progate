package main

import "fmt"

func main() {
	name := "Yahya"
	fmt.Println(name)
	// Panggil function `changeName` dan berikan alamat dari variable `name`
	changeName(&name)
	
	// Cetak variable `name`
	fmt.Println(name)

}

// Tambahkan sebuah parameter pointer dengan tipe *string
func changeName(namePtr *string) {
    // Dereferensikan namePtr dan perbarui nilai dari variable `name` menjadi "Murni"
	*namePtr = "Murni"
	
}
