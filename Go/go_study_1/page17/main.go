package main

func main() {
    n := 3
    switch n {
        // Tambahkan sebuah `case` untuk mencetak "Belum beruntung" ketika `n` bernilai 0
        case 0:
            println("Belum beruntung")
        
        // Tambahkan sebuah `case` untuk mencetak "Sedikit beruntung" ketika `n` bernilai 1 atau 2
        case 1,2:
            println("Sedikit beruntung")
        
        // Tambahkan sebuah `case` untuk mencetak "Beruntung" ketika `n` bernilai 3 atau 4
        case 3,4:
            println("Beruntung")
        
        // Tambahkan sebuah `case` untuk mencetak "Sangat beruntung" ketika `n` bernilai 5
        case 5:
            println("Sangat beruntung")
        
    }
}
