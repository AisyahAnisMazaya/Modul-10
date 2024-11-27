package main

import (
	"fmt"
)

func cariMinMax(kelinci [1000]float64, n int) (float64, float64) {
	min := kelinci[0]
	max := kelinci[0]

	for i := 1; i < n; i++ {
		if kelinci[i] < min {
			min = kelinci[i]
		}
		if kelinci[i] > max {
			max = kelinci[i]
		}
	}

	return min, max
}

func main() {
	var kelinci [1000]float64
	var n int

	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&n)

	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 hingga 1000.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&kelinci[i])
	}

	min, max := cariMinMax(kelinci, n)

	fmt.Println("\nHasil Pencarian:")
	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}
