package main

import (
	"fmt"
)

func main() {
	votes := make(map[int]int)
	total := 0
	valid := 0

	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		total++
		if n >= 1 && n <= 20 {
			valid++
			votes[n]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", total)
	fmt.Printf("Suara sah: %d\n", valid)

	// Cari ketua (suara terbanyak, nomor terkecil jika seri)
	ketuaNo := -1
	ketuaVote := -1
	for i := 1; i <= 20; i++ {
		if votes[i] > ketuaVote {
			ketuaVote = votes[i]
			ketuaNo = i
		}
	}

	// Cari wakil (suara terbanyak kedua, bukan ketua, nomor terkecil jika seri)
	wakilNo := -1
	wakilVote := -1
	for i := 1; i <= 20; i++ {
		if i == ketuaNo {
			continue
		}
		if votes[i] > wakilVote {
			wakilVote = votes[i]
			wakilNo = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketuaNo)
	fmt.Printf("Wakil ketua: %d\n", wakilNo)
}