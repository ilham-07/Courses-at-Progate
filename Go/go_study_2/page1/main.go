package main 

import "fmt"
import "math/rand"
import "time"

func main() {
    rand.Seed(time.Now().Unix())
    
    for i := 1; i <= 3; i++ {
        fmt.Printf("Angka keberuntungan %d:", i)
        
        number := rand.Intn(6)
        switch number {
            case 0:
                fmt.Println("Kurang beruntung")
            case 1, 2:
                fmt.Println("Sedikit beruntung")
            case 3, 4:
                fmt.Println("Beruntung")
            case 5:
                fmt.Println("Sangat beruntung")
        }
    }
}