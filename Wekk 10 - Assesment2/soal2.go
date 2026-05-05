package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [51]mahasiswa

func nilaiPertama(T arrayMahasiswa, N int, NIM string) int {
	for i := 0; i < N; i++ {
		if T[i].NIM == NIM {
			return T[i].nilai
		}
	}
	return -1
}

func nilaiTerbesar(T arrayMahasiswa, N int, NIM string) int {
	max := -1
	for i := 0; i < N; i++ {
		if T[i].NIM == NIM {
			if T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}
	return max
}

func main() {
	var T arrayMahasiswa
	var N int

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}

	var nimCari string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	pertama := nilaiPertama(T, N, nimCari)
	terbesar := nilaiTerbesar(T, N, nimCari)

	if pertama == -1 {
		fmt.Printf("NIM %s tidak ditemukan\n", nimCari)
	} else {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimCari, pertama)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimCari, terbesar)
	}
}