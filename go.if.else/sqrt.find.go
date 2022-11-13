package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Println("Введите число a: ")
	fmt.Scan(&a)

	var b float64
	fmt.Println("Введите число b: ")
	fmt.Scan(&b)

	var c float64
	fmt.Println("Введите число c: ")
	fmt.Scan(&c)

	D := (b * b) - (4 * a * c)

	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Уравнение имеет 2 корня: " + fmt.Sprintf("%.4f", x1) + " и " + fmt.Sprintf("%.4f", x2))
	} else if D == 0 {
		var x float64

		x = (-b) / (2 * a)

		fmt.Print("Уравнение имеет 1 корень: " + fmt.Sprintf("%.4f", x))
	} else {
		fmt.Println("Уравнение не имеет корней")
	}

	fmt.Scan(&a)
}
