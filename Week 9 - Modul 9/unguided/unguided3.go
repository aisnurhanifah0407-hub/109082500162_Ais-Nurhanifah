package main

import "fmt"

func main() {

	var klubA, klubB string
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	var skorA, skorB int
	var skorKlubA []int
	var skorKlubB []int

	i := 1

	// input pertandingan
	for {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		skorKlubA = append(skorKlubA, skorA)
		skorKlubB = append(skorKlubB, skorB)

		i++
	}

	// tampilkan hasil setelah semua pertandingan selesai
	for j := 0; j < len(skorKlubA); j++ {

		if skorKlubA[j] > skorKlubB[j] {
			fmt.Println("Hasil", j+1, ":", klubA)
		} else if skorKlubB[j] > skorKlubA[j] {
			fmt.Println("Hasil", j+1, ":", klubB)
		} else {
			fmt.Println("Hasil", j+1, ": Draw")
		}

	}

	fmt.Println("Pertandingan selesai")
}