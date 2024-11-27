package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	if n <= 0 {
		fmt.Println("Array kosong")
		return
	}

	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	if n <= 0 {
		fmt.Println("Array kosong")
		return 0.0
	}

	var total float64 = 0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var beratBalita arrBalita
	var n int
	var bMin, bMax, rataRata float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Berat balita ke-%d: ", i+1)
		fmt.Scanln(&beratBalita[i])
	}

	hitungMinMax(beratBalita, n, &bMin, &bMax)

	rataRata = rerata(beratBalita, n)

	fmt.Printf("Berat minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat: %.2f kg\n", rataRata)
}
