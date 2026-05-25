package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func insertionSort(p *tabPartai, n int) {
	for i := 1; i < n; i++ {
		temp := p[i]
		j := i - 1
		for j >= 0 && p[j].suara < temp.suara {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = temp
	}
}

func main() {
	var p tabPartai
	var n int = 0
	var x int

	fmt.Print("Masukan proses input suara :\n")
	fmt.Scan(&x)
	for x != -1 {
		idx := posisi(p, n, x)
		if idx == -1 {
			p[n].nama = x
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}
		fmt.Scan(&x)
	}

	insertionSort(&p, n)

	fmt.Print("\nHasil Perhitungan suara :\n")
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d(%d)", p[i].nama, p[i].suara)
	}
	fmt.Println()
}