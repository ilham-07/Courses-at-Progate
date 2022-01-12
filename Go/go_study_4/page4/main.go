package main

import "fmt"

func main() {
	name := "Yahya"

	fmt.Println(name)
	// Deklarasikan pointer `namePtr` dan tetapkan alamat dari variable `name`
	var namePtr *string = &name
    
	// Perbarui nilai dari variable name dengan "Murni" menggunakan dereferensi `namePtr`
	*namePtr = "Murni"
    
	// Mencetak nilai terbaru dari variable `name`
	fmt.Println(name)
	
}
