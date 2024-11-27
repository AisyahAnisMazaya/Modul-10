package main

import (
	"fmt"
)

type Wadah struct {
	berat    [1000]float64
	jumlah   int
	rataRata float64
}

func hitungTotalBerat(w *Wadah, ikan [1000]float64, x, y int) {
	totalBerat := 0.0

	for i := 0; i < x; i += y {
		wadahBerat := 0.0
		for j := i; j < i+y && j < x; j++ {
			wadahBerat += ikan[j]
		}
		w.berat[w.jumlah] = wadahBerat
		w.jumlah++
		totalBerat += wadahBerat
	}

	w.rataRata = totalBerat / float64(w.jumlah)
}

func main() {
	var x, y int
	var ikan [1000]float64
	var wadahIkan = Wadah{}

	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	if x > 1000 {
		fmt.Println("Jumlah ikan tidak boleh lebih dari 1000.")
		return
	}

	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	hitungTotalBerat(&wadahIkan, ikan, x, y)

	fmt.Println("\nTotal berat di setiap wadah:")
	for i := 0; i < wadahIkan.jumlah; i++ {
		fmt.Printf("Wadah %d: %.2f\n", i+1, wadahIkan.berat[i])
	}

	fmt.Printf("\nRata-rata berat per wadah: %.2f\n", wadahIkan.rataRata)
}
