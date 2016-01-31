package main

import (
	"fmt"
	"math"
)

func main() {

	for i := 1.0; i <= 100; i++ {

		if math.Mod(i, 3) == 0 && math.Mod(i, 5) == 0 {
			fmt.Println("FizzBuzz")
		} else if math.Mod(i, 3) == 0 {
			fmt.Println("Fizz")
		} else if math.Mod(i, 5) == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
}
