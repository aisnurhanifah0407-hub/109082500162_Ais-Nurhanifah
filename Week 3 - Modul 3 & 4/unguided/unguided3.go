package main

import "fmt"


func cetakDeret(n int) {

	for n != 1 {
		fmt.Printf("%d ", n) 
		
		if n % 2 == 0 {
		
			n = n / 2
		} else {
		
			n = (n * 3) + 1
		}
	}
	
	fmt.Println(1)
}

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&bilangan)

	cetakDeret(bilangan)
}