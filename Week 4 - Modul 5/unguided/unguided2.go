package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cariFaktor(n, 1)
	fmt.Println()
}

func cariFaktor(n int, i int) {

	if i <= n {
	
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
		
		cariFaktor(n, i+1)
	}
}