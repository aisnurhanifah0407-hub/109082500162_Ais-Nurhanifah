package main

import "fmt"

func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			var x int
			fmt.Scan(&x)
			if x%2 != 0 {
				ganjil = append(ganjil, x)
			} else {
				genap = append(genap, x)
			}
		}

		// Ganjil urut membesar (ascending)
		selectionSortAsc(ganjil)
		// Genap urut mengecil (descending)
		selectionSortDesc(genap)

		// Gabungkan dan cetak: ganjil dulu, lalu genap
		hasil := append(ganjil, genap...)
		for j, v := range hasil {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
}