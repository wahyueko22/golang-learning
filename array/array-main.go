package main

import (
	"fmt"
	"log"

	helper "github.com/wahyueko22/golang-learning/array/modules"
)

func main() {
	fmt.Println("ok")
	//assign element value to array
	var arr [2]int
	arr[0] = 3
	arr[1] = 2
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("\n---------------Example 2--------------------\n")
	for index, element := range arr {
		fmt.Println(index, "=>", element)

	}

	//assign literal array
	var arrStr []string = []string{"one", "two"}
	fmt.Println(arrStr)

	log.Println("main started")
	helper.ForLooping()
	helper.CallFun("paijo")
	helper.GetValue()
	helper.Recursive(2, 5)
}
