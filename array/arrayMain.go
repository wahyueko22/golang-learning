package main

import (
	"fmt"
	"log"

	helper "github.com/wahyueko22/golang-learning/array/modules"
)

func main()  {
	fmt.Println("ok")
	var arr [2]int;
	arr[0] = 3
	arr[1] = 2
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	log.Println("main started")

	helper.ForLooping()
	helper.CallFun("paijo")
	helper.GetValue()
	helper.Recursive(2,5)
}

