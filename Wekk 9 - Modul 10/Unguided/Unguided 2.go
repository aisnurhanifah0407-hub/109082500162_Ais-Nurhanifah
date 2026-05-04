package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64

	fmt.Print("Masukkan jumlah ikan dan kapasitas per wadah: ")
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	var totalWadah []float64

	for i := 0; i < x; i += y {
		var jumlah float64 = 0

		for j := i; j < i+y && j < x; j++ {
			jumlah += ikan[j]
		}

		totalWadah = append(totalWadah, jumlah)
	}

	fmt.Println("Total berat tiap wadah:")
	for i := 0; i < len(totalWadah); i++ {
		fmt.Printf("%.2f ", totalWadah[i])
	}
	fmt.Println()

	var total float64 = 0
	for i := 0; i < len(totalWadah); i++ {
		total += totalWadah[i]
	}

	rata := total / float64(len(totalWadah))

	fmt.Printf("Rata-rata berat per wadah: %.2f\n", rata)
}