package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan informasi teman
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Data teman-teman
var temanTeman = map[int]Teman{
	1:  {"Vino", "Jakarta", "Software Engineer", "Saya suka dengan kecepatan dan kemudahan penggunaan Go."},
	2:  {"Hapi", "Bandung", "Data Scientist", "Go memiliki banyak library yang cocok untuk kebutuhan analisis data saya."},
	3:  {"Mauladi", "Surabaya", "Web Developer", "Go sangat efisien untuk membangun aplikasi web yang scalable."},
	4:  {"Nadzar", "Yogyakarta", "Mobile App Developer", "Saya tertarik dengan performa tinggi Go dalam membangun aplikasi mobile."},
	5:  {"Ahmad Rahman", "Semarang", "DevOps Engineer", "Go merupakan pilihan utama dalam pengembangan infrastruktur berbasis cloud."},
	6:  {"Siti Aisyah", "Medan", "Blockchain Developer", "Saya ingin mengembangkan aplikasi blockchain dengan Go karena keamanannya."},
	7:  {"Budi Santoso", "Makassar", "Game Developer", "Go sangat cocok untuk pengembangan game berbasis server."},
	8:  {"Dewi Susanti", "Palembang", "UI/UX Designer", "Saya ingin mempelajari Go untuk memahami teknologi yang digunakan dalam produk digital."},
	9:  {"Rudi Hartono", "Denpasar", "Embedded Systems Engineer", "Go dapat digunakan untuk pengembangan perangkat embedded dengan performa tinggi."},
	10: {"Lina Wijaya", "Bekasi", "Machine Learning Engineer", "Saya tertarik dengan kemampuan Go dalam pengolahan data untuk machine learning."},
}

// Function untuk menampilkan biodata teman berdasarkan nomor absen
func TampilkanBiodata(nomorAbsen int) {
	teman, ok := temanTeman[nomorAbsen]
	if !ok {
		fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan.")
		return
	}

	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}

func main() {
	// Cek jumlah argumen yang diberikan
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run biodata.go [nomor absen]")
		return
	}

	// Ambil nomor absen dari argumen kedua
	nomorAbsen := os.Args[1]

	// Konversi nomor absen menjadi integer
	var nomorAbsenInt int
	_, err := fmt.Sscanf(nomorAbsen, "%d", &nomorAbsenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa bilangan bulat.")
		return
	}

	// Tampilkan biodata teman berdasarkan nomor absen
	TampilkanBiodata(nomorAbsenInt)
}
