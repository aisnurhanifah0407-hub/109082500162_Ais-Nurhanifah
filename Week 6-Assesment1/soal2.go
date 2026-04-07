package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
	} else if tujuan == "mancanegara" {
		if jumlahHari > 8 {
			return 8
		}
	}
	return jumlahHari
}

func biayaPerHari(jumlahMhs int) int {
	return 620000 * jumlahMhs
}

func perhitunganBiaya(jumlahMhs, lama int, tujuan string, totalBiaya *float64) {
	hari := tanggunganHari(lama, tujuan)
	baseBiaya := biayaPerHari(jumlahMhs)

	if tujuan == "domestik" {
		*totalBiaya = float64(hari * baseBiaya)
	} else if tujuan == "mancanegara" {
		*totalBiaya = float64(hari) * (float64(baseBiaya) * 1.5)
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}