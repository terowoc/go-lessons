package main

import "fmt"

// Hello World in Go
// func main() {
// 	fmt.Println("Hello, World!")
// }

func main() {
	var name string
	var age int8

	fmt.Println("What is your name?")
	fmt.Scan(&name)

	fmt.Println("How old are you?")
	fmt.Scan(&age)

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}