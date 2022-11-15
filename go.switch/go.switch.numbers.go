package main

import "fmt"

func main() {
	num := 10

	switch {
	case num > 0:
		fmt.Println("Number is positive")
		fallthrough //используется для проверки последующих за этим кейсом условий
	case num < 11:
		fmt.Println("num < 11")
	default:
		fmt.Println("Number equival zero or negative")
	}
}
