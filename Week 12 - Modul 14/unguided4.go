package main

import (
	"bufio"
	"fmt"
	"os"
)

const nMax = 7919

type Buku struct {
	id       string
	judul    string
	penulis  string
	penerbit string
	eksemplar int
	tahun    int
	rating   int
}

type DaftarBuku [nMax]Buku

var (
	pustaka  DaftarBuku
	nPustaka int
	reader   = bufio.NewReader(os.Stdin)
)

func bacaLine() string {
	line, _ := reader.ReadString('\n')
	// trim \r\n
	for len(line) > 0 && (line[len(line)-1] == '\n' || line[len(line)-1] == '\r') {
		line = line[:len(line)-1]
	}
	return line
}

// DaftarkanBuku: membaca n data buku dari input
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Print("ID Buku       : ")
		pustaka[i].id = bacaLine()

		fmt.Print("Judul         : ")
		pustaka[i].judul = bacaLine()

		fmt.Print("Penulis       : ")
		pustaka[i].penulis = bacaLine()

		fmt.Print("Penerbit      : ")
		pustaka[i].penerbit = bacaLine()

		fmt.Print("Eksemplar     : ")
		fmt.Fscan(reader, &pustaka[i].eksemplar)
		reader.ReadString('\n')

		fmt.Print("Tahun Terbit  : ")
		fmt.Fscan(reader, &pustaka[i].tahun)
		reader.ReadString('\n')

		fmt.Print("Rating        : ")
		fmt.Fscan(reader, &pustaka[i].rating)
		reader.ReadString('\n')

		fmt.Println()
	}
}

// CetakTerfavorit: cetak buku dengan rating tertinggi (belum diurutkan)
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}
	maxIdx := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[maxIdx].rating {
			maxIdx = i
		}
	}
	b := pustaka[maxIdx]
	fmt.Println("=== Buku Terfavorit ===")
	fmt.Println("Judul    :", b.judul)
	fmt.Println("Penulis  :", b.penulis)
	fmt.Println("Penerbit :", b.penerbit)
	fmt.Println("Tahun    :", b.tahun)
	fmt.Println("Rating   :", b.rating)
}

// UrutBuku: insertion sort descending berdasarkan rating
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

// Cetak5Terbaru: cetak 5 buku dengan rating tertinggi (sudah diurutkan)
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	fmt.Println("=== 5 Buku dengan Rating Tertinggi ===")
	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
}

// CariBuku: binary search berdasarkan rating (array sudah diurutkan descending)
// Binary search pada array descending: cari dari kiri, nilai besar di kiri
func CariBuku(pustaka DaftarBuku, n int, r int) {
	low := 0
	high := n - 1
	hasil := -1

	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			hasil = mid
			break
		} else if pustaka[mid].rating > r {
			// nilai mid lebih besar, target ada di kanan (index lebih besar)
			low = mid + 1
		} else {
			// nilai mid lebih kecil, target ada di kiri (index lebih kecil)
			high = mid - 1
		}
	}

	if hasil == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		b := pustaka[hasil]
		fmt.Println("=== Data Buku Ditemukan ===")
		fmt.Println("Judul     :", b.judul)
		fmt.Println("Penulis   :", b.penulis)
		fmt.Println("Penerbit  :", b.penerbit)
		fmt.Println("Tahun     :", b.tahun)
		fmt.Println("Eksemplar :", b.eksemplar)
		fmt.Println("Rating    :", b.rating)
	}
}

func main() {
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Fscan(reader, &nPustaka)
	reader.ReadString('\n')
	fmt.Println()

	// 1. Daftarkan buku
	DaftarkanBuku(&pustaka, nPustaka)

	// 2. Cetak buku terfavorit (sebelum diurutkan)
	CetakTerfavorit(pustaka, nPustaka)
	fmt.Println()

	// 3. Urutkan buku berdasarkan rating (descending)
	UrutBuku(&pustaka, nPustaka)

	// 4. Cetak 5 buku terbaru (rating tertinggi)
	Cetak5Terbaru(pustaka, nPustaka)
	fmt.Println()

	// 5. Cari buku berdasarkan rating
	var ratingCari int
	fmt.Print("Masukkan rating yang dicari: ")
	fmt.Fscan(reader, &ratingCari)
	CariBuku(pustaka, nPustaka, ratingCari)
}