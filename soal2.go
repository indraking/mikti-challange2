package main

import "fmt"

// Fungsi menghitung BMI
func bmi(massa float64, tinggi float64) float64 {
	bmi := massa / (tinggi * tinggi)
	return bmi
}

func main() {

	mark1 := bmi(78, 1.69)
	john1 := bmi(92, 1.95)
	// Variabel ini untuk mengecek true atau false
	markHigherBMI1 := mark1 > john1
	fmt.Printf("Hasil BMI mark 1: %.2f \n", mark1)
	fmt.Printf("Hasil BMI john 1: %.2f \n", john1)
	fmt.Println("Mark memiliki BMI lebih tinggi dari John : ", markHigherBMI1)

	mark2 := bmi(95, 1.88)
	john2 := bmi(85, 1.76)
	// Variabel ini untuk mengecek true atau false
	markHigherBMI2 := mark2 > john2
	fmt.Printf("Hasil BMI mark 1: %.2f \n", mark2)
	fmt.Printf("Hasil BMI john 1: %.2f \n", john2)
	fmt.Println("Mark memiliki BMI lebih tinggi dari John : ", markHigherBMI2)

}
