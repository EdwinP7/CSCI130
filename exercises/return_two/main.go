package main

import (
	"fmt"
)

func main() {
	fmt.Println(byTwo(1))
	fmt.Println(byTwo(2))

	half := func(i int) (int, bool) {
		return i / 2, i%2 == 0
	}
	
	fmt.Println(half(1))
	fmt.Println(half(2))

}

func byTwo(i int) (int, bool) {

	return i / 2, i%2 == 0
}
