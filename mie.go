package main

type Rasa struct {
    Nama  string
    Harga int
}

type Pesanan struct {
    NomorRasa int
    Jumlah    int
}

type GerobakMie struct {
    DaftarRasa []Rasa
    DaftarPesan []Pesanan
}

func (g *GerobakMie) TambahRasa(nama string, harga int) {
    rasaBaru := Rasa{Nama: nama, Harga: harga}
    g.DaftarRasa = append(g.DaftarRasa, rasaBaru)
}

func (g *GerobakMie) TampilkanMenu() {
    fmt.Println("Menu Mie Ayam:")
    for i, rasa := range g.DaftarRasa {
        nomor := i + 1
        fmt.Printf("%d. %s - Rp %d\n", nomor, rasa.Nama, rasa.Harga)
    }
}

func (g *GerobakMie) Pesan(p Pesanan) string {
    if p.NomorRasa < 1 || p.NomorRasa > len(g.DaftarRasa) {
        return "Nomor rasa tidak valid"
    }
    
    if p.Jumlah <= 0 {
        return "Jumlah harus lebih dari 0"
    }
    
    g.DaftarPesan = append(g.DaftarPesan, p)
    rasa := g.DaftarRasa[p.NomorRasa-1]
    total := rasa.Harga * p.Jumlah
    
    return fmt.Sprintf("Berhasil memesan %d porsi %s = Rp %d", 
        p.Jumlah, rasa.Nama, total)
}

func (g *GerobakMie) HitungTotal() string {
    if len(g.DaftarPesan) == 0 {
        return "Belum ada pesanan"
    }
    
    var totalSemua int
    var detail strings.Builder
    
    detail.WriteString("Rincian Pesanan:\n")
    
    for _, pesan := range g.DaftarPesan {
        rasa := g.DaftarRasa[pesan.NomorRasa-1]
        subtotal := rasa.Harga * pesan.Jumlah
        totalSemua += subtotal
        
        detail.WriteString(fmt.Sprintf("- %s: %d x Rp %d = Rp %d\n", 
            rasa.Nama, pesan.Jumlah, rasa.Harga, subtotal))
    }
    
    detail.WriteString(fmt.Sprintf("\nTOTAL: Rp %d", totalSemua))
    
    return detail.String()
}
