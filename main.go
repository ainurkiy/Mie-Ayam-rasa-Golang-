package main

import (
    "fmt"
    "strings"
)

func main() {
    gerobak := GerobakMie{}
    gerobak.TambahRasa("Original", 10000)
    gerobak.TambahRasa("Pedas", 12000)
    gerobak.TambahRasa("Soto", 15000)
    
    fmt.Println("=== PROJECT MIE AYAM RASA GOLANG ===")
    gerobak.TampilkanMenu()
    
    for {
        fmt.Print("\nPilih nomor rasa (0 untuk selesai): ")
        var pilihan int
        fmt.Scanln(&pilihan)
        
        if pilihan == 0 {
            break
        }
        
        fmt.Print("Jumlah porsi: ")
        var jumlah int
        fmt.Scanln(&jumlah)
        
        pesan := Pesanan{NomorRasa: pilihan, Jumlah: jumlah}
        hasil := gerobak.Pesan(pesan)
        fmt.Println(hasil)
    }
    
    total := gerobak.HitungTotal()
    fmt.Printf("\n=== TOTAL PEMBELIAN ===\n%s\n", total)
}
