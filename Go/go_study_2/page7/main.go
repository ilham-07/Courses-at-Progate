package main

import "fmt"
// Lakukan import paket `math/rand`
import "math/rand"

func main() {
    
    for i := 1; i <= 5; i++ {
        // Hasilkan dan cetak sebuah angka integer dari 0 hingga 9
        fmt.Println(rand.Intn(10))
    }
}
