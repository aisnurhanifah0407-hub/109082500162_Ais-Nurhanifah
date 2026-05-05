package main

import "fmt"

const nProv int = 10

type NamaProv [10]string
type PopProv [10]int
type TumbuhProv [10]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIdx := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := (tumbuh[i] + 1) * float64(pop[i])
			fmt.Printf("%s %.0f\n", prov[i], prediksi)
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv

	InputData(&prov, &pop, &tumbuh)

	var namaCari string
	fmt.Scan(&namaCari)

	idxTercepat := ProvinsiTercepat(tumbuh)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", prov[idxTercepat])

	idxCari := IndeksProvinsi(prov, namaCari)
	if idxCari == -1 {
		fmt.Printf("\nData provinsi yang dicari : %s tidak ditemukan\n", namaCari)
	} else {
		fmt.Printf("\nData provinsi yang dicari : %s\n", prov[idxCari])
	}

	fmt.Println()
	Prediksi(prov, pop, tumbuh)
}
