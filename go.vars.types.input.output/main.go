package main

import "fmt"

func main() {
	var age int8 = 23
	fmt.Println(age)

	var number float64 = 275.7867
	fmt.Println(number)

	age0 := 16
	fmt.Println(age0)

	var name string = "Tagilla"
	fmt.Println(name)
	fmt.Println(len(name))

	fmt.Scanln(&name)
	fmt.Println("Hello, " + name + "!")

	fmt.Println("How old are you?")
	fmt.Scan(&age)
	fmt.Println("Your age, " + fmt.Sprint(age) + " years old!")
}
