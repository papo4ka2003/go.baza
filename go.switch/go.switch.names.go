package main

import "fmt"

func main() {
	name := "Forrest"

	switch name {
	case "Forrest":
		fmt.Println("Hello Forrest!")
	case "Julia":
		fmt.Println("Hello Julia!")
	default:
		fmt.Println("I dont know you!")
	}
}
