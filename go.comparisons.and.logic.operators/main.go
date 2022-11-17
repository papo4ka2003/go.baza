package main

import "fmt"

func main() {
	a := 1

	if a >= 1 {
		fmt.Println("a > 1")
	}

	a = 3
	b := 10

	if a > 1 && b > 8 { //  1 && 1 == 1
		fmt.Println("TRUE!!!") //  1 && 0 == 0
	} //  0 && 1 == 0
	//  0 && 0 == 0

	if a > 1 || b > 8 { //  1 || 1 == 1
		fmt.Println("TRUE!!!") //  1 || 0 == 1
	} //  0 || 1 == 1
	//  0 || 0 == 0
	if a != 5 {
		fmt.Println("n != 5")
	}
}
