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

		x1 = (-b + math.Sqrt(D)) / 2 * a
		x2 = (-b - math.Sqrt(D)) / 2 * a

		fmt.Println("Уравнение имеет 2 корня: " + fmt.Sprint(x1) + " и " + fmt.Sprint(x2))
	} else if D == 0 {
		var x1 float64

		x1 = (-b - math.Sqrt(D)) / 2 * a

		fmt.Println("Уравнение имеет 1 корень: " + fmt.Sprint(x1))
	} else {
		fmt.Println("Уравнение не имеет корней")
	}

}
