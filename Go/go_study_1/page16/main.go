package main

func main() {
    x := 20
    // Ketika x lebih besar dari atau sama dengan 10 dan kurang dari atau sama dengan 30,
    // cetak pesan "x paling sedikit 10 dan paling banyak 30"
    if x >= 10 && x <= 30 {
        println("x paling sedikit 10 dan paling banyak 30")
    }
    
    y := 60
    // Ketika y kurang dari 10 atau lebih besar dari 30,
    // cetak pesan "y kurang dari 10 atau lebih besar dari 30"
    if y < 10 || y > 30 {
        println("y kurang dari 10 atau lebih besar dari 30")
    }
    
    z := 55
    // Ketika z tidak sama dengan 77,
    // cetak pesan "z tidak sama dengan 77"
    if !(z == 77) {
        println("z tidak sama dengan 77")
    }
    
}
