package main

import (
	"fmt"
	"math"
)

func main() {

	sumOfSquares := 0.0
	sum := 0.0
	squareOfSum := 0.0

	for i := 1.0; i <= 100.0; i++ {
		sumOfSquares += math.Pow(i, 2)
		sum += i
	}

	squareOfSum = math.Pow(sum, 2)
	difference := squareOfSum - sumOfSquares
	fmt.Println(difference)

}
