package main

import "fmt"
import "math/rand"
// Lakukan import paket `time`
import "time"

func main() {
    // Berikan nilai benih(seed) pada penghasil nomor acak
    rand.Seed(time.Now().Unix())
    for i := 1; i <= 5; i++ {
        fmt.Println(rand.Intn(10))
    }
}
