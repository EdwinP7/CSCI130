package main

import "fmt"

func main() {

	sum(1, 6, 10, 1000, 99, 101, 55, 10)
	fmt.Println((true && false) || (false && true) || (!(false && false)))

	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()

}

func sum(nums ...int) int {

	fmt.Println(nums, " ")
	greatest := 0
	for _, num := range nums {
		if num > greatest {
			greatest = num
		}
	}
	return greatest
}

func foo(numbers ...int) {

	fmt.Println("function foo: ")
	for _, num := range numbers {
		fmt.Println(num)
	}
}
