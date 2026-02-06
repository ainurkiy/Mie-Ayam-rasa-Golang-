
# 🍜 Mie Ayam Rasa Golang

Projek simulasi gerobak mie ayam dengan tiga varian rasa, dibangun menggunakan bahasa pemrograman Go (Golang).

## 📋 Fitur

- **3 Varian Rasa**:
  - Original (Rp 10,000)
  - Pedas (Rp 12,000)
  - Soto (Rp 15,000)
- **Sistem Pemesanan Interaktif**
- **Perhitungan Total Otomatis**
- **Struktur Kode Modular**
- **Validasi Input**

## 🚀 Instalasi dan Menjalankan

### Prasyarat
- Go 1.20 atau lebih tinggi

### Clone Repository
```bash
git clone https://ainurkiy/Mie-Ayam-rasa-Golang.git
cd Mie-Ayam-rasa-Golang
```

Menjalankan Program

```bash
# Cara 1: Menggunakan go run
go run main.go mie.go

# Cara 2: Build dan jalankan
go build -o mieayam
./mieayam
```

📁 Struktur File

```
Mie-Ayam-rasa-Golang/
├── main.go      # Entry point program
├── mie.go       # Logika bisnis mie ayam
├── go.mod       # Module configuration
└── README.md    # Dokumentasi
```

🎮 Cara Menggunakan

1. Program akan menampilkan menu tiga rasa mie ayam
2. Pilih nomor rasa (1-3)
3. Masukkan jumlah porsi
4. Ulangi sampai selesai (ketik 0 untuk keluar)
5. Program akan menampilkan total pembelian

Contoh Interaksi

```
=== PROJECT MIE AYAM RASA GOLANG ===
Menu Mie Ayam:
1. Original - Rp 10000
2. Pedas - Rp 12000
3. Soto - Rp 15000

Pilih nomor rasa (0 untuk selesai): 1
Jumlah porsi: 2
Berhasil memesan 2 porsi Original = Rp 20000

Pilih nomor rasa (0 untuk selesai): 0

=== TOTAL PEMBELIAN ===
Rincian Pesanan:
- Original: 2 x Rp 10000 = Rp 20000

TOTAL: Rp 20000
```

🏗️ Arsitektur Kode

main.go

· Entry point aplikasi
· Menampilkan antarmuka pengguna
· Menangani input dari pengguna

mie.go

· Struct Rasa: Menyimpan informasi varian mie
· Struct Pesanan: Menyimpan data pemesanan
· Struct GerobakMie: Logika bisnis utama
  · TambahRasa(): Menambah varian rasa
  · TampilkanMenu(): Menampilkan daftar menu
  · Pesan(): Memproses pemesanan
  · HitungTotal(): Menghitung total pembelian

📝 Contoh Kode

```go
// Menambahkan rasa baru
gerobak.TambahRasa("Original", 10000)

// Memesan mie
pesan := Pesanan{NomorRasa: 1, Jumlah: 2}
hasil := gerobak.Pesan(pesan)
```

🤝 Kontribusi

1. Fork repository
2. Buat branch fitur (git checkout -b fitur-baru)
3. Commit perubahan (git commit -m 'Menambah fitur baru')
4. Push ke branch (git push origin fitur-baru)
5. Buat Pull Request

📄 Lisensi

Proyek ini dilisensikan di bawah MIT License.

✨ Kontributor

· ainurkiy

---

⭐ Jangan lupa kasih bintang jika proyek ini bermanfaat!

```
