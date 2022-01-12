package main

import "fmt"

func main() {
	name := "Yahya"

	fmt.Println(name)
	fmt.Println(&name)
	// Deklarasikan pointer `namePtr` bertipe string dan tetapkan alamat dari variable `name`
	var namePtr *string = &name
	
	// Cetak nilai dari `namePtr`
	fmt.Println(namePtr)

}
