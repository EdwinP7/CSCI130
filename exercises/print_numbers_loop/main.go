package main

import (
	"fmt"
	"math"
)

func main() {

	for i := 0.0; i <= 100.0; i++ {
		if math.Mod(i, 2) == 0 {
			fmt.Println(i)
		}
	}

}
