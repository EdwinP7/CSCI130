package main

import (
	"fmt"
	"math"
)

func main() {

	sum := 0.0

	for i := 0.0; i < 1000; i++ {

		if math.Mod(i, 3) == 0 {
			sum += i
		} else if math.Mod(i, 5) == 0 {
			sum += i
		}

	}

	fmt.Println(sum)
}
