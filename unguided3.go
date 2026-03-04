package main

import "fmt"

func main() {
	var totalGram int
	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&totalGram)

	kg, gr := totalGram/1000, totalGram%1000
	biayaKg := kg * 10000
	biayaGr := 0

	if kg < 10 {
		if gr >= 500 {
			biayaGr = gr * 5  
		} else {
			biayaGr = gr * 15 
		}
	}

	fmt.Println("===== Detail Perhitungan =====")
	fmt.Printf("Detail berat : %d kg + %d gram\n", kg, gr)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayaKg, biayaGr)
	fmt.Printf("Total biaya: Rp %d\n", biayaKg + biayaGr)
}