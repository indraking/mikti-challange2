package main

import "fmt"

// Fungsi menghitung rata-rata
func calculateAvg(numbers ...int) float64 {
	result := 0
	for _, number := range numbers {
		result += number
	}
	avg := float64(result) / float64(len(numbers))
	return avg
}

func main() {

	dolphin1 := calculateAvg(96, 108, 89)
	koala1 := calculateAvg(88, 91, 110)
	dolphinBonus1 := calculateAvg(97, 112, 101)
	koalaBonus1 := calculateAvg(109, 95, 123)
	dolphinBonus2 := calculateAvg(97, 112, 101)
	koalaBonus2 := calculateAvg(109, 95, 106)
	minimumScore := 100

	// Pengondisian untuk menentukan hasil yang tepat
	if dolphin1 < koala1 {
		fmt.Printf("Pemenangnya Koala: %.2f \n", koala1)
	} else if dolphin1 > koala1 {
		fmt.Printf("Pemenangnya Dolphin : %.2f \n", dolphin1)
	} else {
		fmt.Printf("Hasil Seri tidak ada yang menang \n")
	}

	if dolphinBonus1 > float64(minimumScore) && koalaBonus1 > float64(minimumScore) {
		if dolphinBonus1 < koalaBonus1 {
			fmt.Printf("Pemenangnya Koala: %.2f \n", koalaBonus1)
		} else if dolphinBonus1 > koalaBonus1 {
			fmt.Printf("Pemenangnya Dolphin : %.2f \n", dolphinBonus1)
		}
	} else {
		fmt.Printf("Hasil Seri tidak ada yang menang karena score lebih kecil 100 \n")
	}

	if dolphinBonus2 >= float64(minimumScore) && koalaBonus2 >= float64(minimumScore) {
		if dolphinBonus2 < koalaBonus2 {
			fmt.Printf("Pemenangnya Koala: %.2f \n", koalaBonus2)
		} else if dolphinBonus2 > koalaBonus2 {
			fmt.Printf("Pemenangnya Dolphin : %.2f \n", dolphinBonus2)
		} else {
			fmt.Printf("Hasil Seri tidak ada yang menang karena score lebih besar 100 \n")
		}
	} else {
		fmt.Println("Tidak ada tim yang memenangkan trofi")
	}
}
