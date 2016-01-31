package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("What is your name? ")
	name := ""
	fmt.Scanf("%v", &name)
	fmt.Println("Hello there, " + name)

	fmt.Println("Enter two numbers, one larger than the other.")
	var i, j float64
	fmt.Scanf("%v", &i)
	fmt.Scanf("%v", &j)
	z := math.Mod(i, j)
	fmt.Println(z)

}
