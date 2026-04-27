package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	// input array
	for i := 0; i < n; i++ {
		fmt.Print("Elemen ke-", i, ": ")
		fmt.Scan(&arr[i])
	}

	// a. tampilkan seluruh array
	fmt.Println("\nIsi array:")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// b. indeks ganjil
	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 0; i < len(arr); i++ {
		if i%2 == 1 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// c. indeks genap
	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// d. indeks kelipatan x
	var x int
	fmt.Print("\nMasukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// e. hapus elemen pada indeks tertentu
	var idx int
	fmt.Print("\nMasukkan indeks yang akan dihapus: ")
	fmt.Scan(&idx)

	arr = append(arr[:idx], arr[idx+1:]...)

	fmt.Println("Array setelah penghapusan:")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// f. rata-rata
	sum := 0
	for _, v := range arr {
		sum += v
	}

	mean := float64(sum) / float64(len(arr))
	fmt.Println("\nRata-rata:", mean)

	// g. standar deviasi
	var total float64
	for _, v := range arr {
		total += math.Pow(float64(v)-mean, 2)
	}

	std := math.Sqrt(total / float64(len(arr)))
	fmt.Println("Standar deviasi:", std)

	// h. frekuensi suatu bilangan
	var cari int
	fmt.Print("\nMasukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)

	freq := 0
	for _, v := range arr {
		if v == cari {
			freq++
		}
	}

	fmt.Println("Frekuensi", cari, "=", freq)
}