package main

import "fmt"

type suhu float64

func CelciusToReamur(Celcius suhu) suhu {
	return Celcius * 4 / 5
}

func CelciusToFahrenheit(Celcius suhu) suhu {
	return (Celcius * 9 / 5) + 32
}

func CelciusToKelvin(Celcius suhu) suhu {
	return Celcius + 273.15
}

func main() {
	var Celcius suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&Celcius)

	fmt.Println()

	fmt.Println(Celcius, "celcius =", CelciusToReamur(Celcius), "reamur")
	fmt.Println(Celcius, "celcius =", CelciusToFahrenheit(Celcius), "fahrenheit")
	fmt.Println(Celcius, "celcius =", CelciusToKelvin(Celcius), "kelvin")
}