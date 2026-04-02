package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	tampilkanGanjil(n, 1)
	fmt.Println()
}

func tampilkanGanjil(n int, i int) {
	if i <= n {

		if i%2 != 0 {
			fmt.Printf("%d ", i)
		}
		tampilkanGanjil(n, i+1)
	}
}