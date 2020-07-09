package main

import (
	"fmt"
)

func main() {
	sum(1, 2, 3, 4, 5)
	var test = float64(100 / 3)
	fmt.Println(test)
}

func sum(numbers ...int) {
	var total = 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println("total : ", total)
	return
}
