package main

import (
	"fmt"
	"math"
)

type titik struct {
	x int
	y int
}

type lingkaran struct {
	pusat  titik
	radius int
}

// fungsi menghitung jarak dua titik
func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

// fungsi mengecek apakah titik berada di dalam lingkaran
func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {

	var c1, c2 lingkaran
	var p titik

	// input lingkaran 1
	fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.radius)

	// input lingkaran 2
	fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.radius)

	// input titik
	fmt.Scan(&p.x, &p.y)

	in1 := didalam(c1, p)
	in2 := didalam(c2, p)

	if in1 && in2 {
		fmt.Println("Titik didalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik didalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik didalam lingkaran 2")
	} else {
		fmt.Println("Titik diluar lingkaran 1 dan 2")
	}
}