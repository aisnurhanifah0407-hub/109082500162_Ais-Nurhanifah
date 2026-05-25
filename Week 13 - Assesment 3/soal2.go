package main

import "fmt"

const NMAX = 1001

type Pemain struct {
	nama   string
	gol    int
	assist int
}

type arrPemain [NMAX]Pemain

func SelectionSort(T *arrPemain, n int) {
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if T[j].gol > T[maxIdx].gol {
				maxIdx = j
			} else if T[j].gol == T[maxIdx].gol && T[j].assist > T[maxIdx].assist {
				maxIdx = j
			}
		}
		T[i], T[maxIdx] = T[maxIdx], T[i]
	}
}

func main() {
	var A arrPemain
	var n int
	var firstName, lastName string

	fmt.Print("Masukan Data Input :\n")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&firstName, &lastName, &A[i].gol, &A[i].assist)
		A[i].nama = firstName + " " + lastName
	}

	SelectionSort(&A, n)

	fmt.Print("\nHasil Sorting :\n")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %d\n", A[i].nama, A[i].gol, A[i].assist)
	}
}