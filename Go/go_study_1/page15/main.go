package main

func main() {
    money := 200
    price := 200
    
    // Tambahkan pernyataan `else if` dan  `else` berikut ini
    if money > price {
        println("Anda bisa membeli item ini")
    } else if money == price {
        println("Anda bisa membeli item ini, namun uang tidak akan tersisa")
    } else {
        println("Anda tidak memiliki cukup uang")
    }
}
