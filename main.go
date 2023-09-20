package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Struct untuk menyimpan informasi barang
type Item struct {
	Nama  string
	Harga float64
}

// Fungsi utama
func main() {
	// Daftar barang
	daftarBarang := map[string]Item{
		"barang1": {"pulpen", 10.0},
		"barang2": {"buku", 20.0},
		"barang3": {"pensil warna", 30.0},
	}

	// Keranjang belanja
	keranjang := make(map[string]int)

	// Input nama dan jumlah barang yang dibeli
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Daftar Barang:")
		for key, item := range daftarBarang {
			fmt.Printf("%s: %s (Harga: Rp%.2f)\n", key, item.Nama, item.Harga)
		}

		fmt.Print("Masukkan kode barang yang dibeli (q untuk keluar): ")
		scanner.Scan()
		input := strings.ToLower(scanner.Text())

		if input == "q" {
			break
		}

		item, found := daftarBarang[input]
		if !found {
			fmt.Println("Barang tidak ditemukan.")
			continue
		}

		fmt.Printf("Masukkan jumlah %s yang dibeli: ", item.Nama)
		scanner.Scan()
		jumlah := 0
		_, err := fmt.Sscanf(scanner.Text(), "%d", &jumlah)
		if err != nil || jumlah <= 0 {
			fmt.Println("Jumlah barang tidak valid.")
			continue
		}

		keranjang[input] += jumlah
	}

	// Hitung total dan total biaya
	var total int
	var totalBiaya float64

	fmt.Println("Invoice:")
	for kode, jumlah := range keranjang {
		item := daftarBarang[kode]
		subtotal := item.Harga * float64(jumlah)
		fmt.Printf("%s (%s) x %d = Rp%.2f\n", item.Nama, kode, jumlah, subtotal)
		total += jumlah
		totalBiaya += subtotal
	}

	fmt.Printf("Total Barang: %d\n", total)
	fmt.Printf("Total Biaya: Rp%.2f\n", totalBiaya)

	// Simpan invoice ke dalam file
	file, err := os.Create("invoice.txt")
	if err != nil {
		fmt.Println("Gagal menyimpan invoice ke file.")
	} else {
		defer file.Close()

		fmt.Fprintf(file, "Invoice:\n")
		for kode, jumlah := range keranjang {
			item := daftarBarang[kode]
			subtotal := item.Harga * float64(jumlah)
			fmt.Fprintf(file, "%s (%s) x %d = Rp%.2f\n", item.Nama, kode, jumlah, subtotal)
		}
		fmt.Fprintf(file, "Total Barang: %d\n", total)
		fmt.Fprintf(file, "Total Biaya: Rp%.2f\n", totalBiaya)

		fmt.Println("Invoice telah disimpan ke invoice.txt")
	}
}
