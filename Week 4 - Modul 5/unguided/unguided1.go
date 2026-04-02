package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan nilai n: ")
    fmt.Scan(&n)

    fmt.Printf("Deret Fibonacci hingga suku ke-%d:\n", n)
    
    for i := 0; i <= n; i++ {
        fmt.Printf("%d ", fibonacci(i))
    }
    fmt.Println() 
}

func fibonacci(n int) int {
    if n <= 1 {
        return n
    } else {
        return fibonacci(n-1) + fibonacci(n-2)
    }
}