package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {

	*n = 0
	var input string
	for *n < NMAX {

		fmt.Scan(&input)
		if input == "." {
			break
		}
	
		t[*n] = rune(input[0])
		(*n)++
	}
}

func cetakArray(t tabel, n int) {
	
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	
	for i := 0; i < n/2; i++ {
		
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	
	var temp tabel = t
	
	balikanArray(&temp, n)
	
	for i := 0; i < n; i++ {
		if t[i] != temp[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks : ")
	
	isiArray(&tab, &m)

	hasilPalindrom := palindrom(tab, m)
	fmt.Printf("Palindrom : %v\n", hasilPalindrom)

	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)
}