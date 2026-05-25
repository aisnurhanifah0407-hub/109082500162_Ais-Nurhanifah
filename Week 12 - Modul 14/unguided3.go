package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var arr []int

	for {
		var x int
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		arr = append(arr, x)
	}

	insertionSort(arr)

	// Cetak array terurut
	for i, v := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()

	// Cek jarak
	if len(arr) < 2 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}

	jarak := arr[1] - arr[0]
	tetap := true
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != jarak {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}