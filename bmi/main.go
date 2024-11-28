package main

import (
	"fmt"
)

func main() {
	// Foydalanuvchining vazni va bo'yi
	var weight, height float64

	// Foydalanuvchidan ma'lumotlarni so'rash
	fmt.Print("Vazningizni kiriting (kg): ")
	fmt.Scan(&weight)

	fmt.Print("Bo'yingizni kiriting (m): ")
	fmt.Scan(&height)

	// BMI hisoblash
	bmi := weight / (height * height)

	// Natijani chiqarish
	fmt.Printf("Sizning BMI: %.5f\n", bmi)

	// BMI bo'yicha tavsiyalar
	if bmi < 18.5 {
		fmt.Println("Sizning vazningiz kam (underweight).")
	} else if bmi >= 18.5 && bmi < 24.9 {
		fmt.Println("Sizning vazningiz normal.")
	} else if bmi >= 25 && bmi < 29.9 {
		fmt.Println("Siz ortiqcha vaznga egasiz (overweight).")
	} else {
		fmt.Println("Sizda semizlik (obesity) bor.")
	}
}
