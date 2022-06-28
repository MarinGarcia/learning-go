package main

import "fmt"

var name string
var age int
var height float32

func main() {
	fmt.Print("Name: ")
	fmt.Scanf("%s", &name)

	fmt.Print("Age: ")
	fmt.Scanf("%d", &age)

	fmt.Print("Height: ")
	fmt.Scanf("%f", &height)

	fmt.Printf("Hellow %s, your age is %d and your height is %.2f", name, age, height)
}
