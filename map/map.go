package main

import "fmt"

func main() {
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	//var chicken5 = *new(map[string]int)
	chicken3["one"] = 1
	chicken4["one"] = 1
	//chicken5["one"] = 1
	fmt.Println("result", chicken3["one"], chicken4["one"])
	mapToAnotherFunction(chicken3)
	for key, val := range chicken3 {
		fmt.Println(key, "  \t:", val)
	}
}

func mapToAnotherFunction(chicken3 map[string]int) {
	chicken3["two"] = 2
	for key, val := range chicken3 {
		fmt.Println(key, "  \t:", val)
	}
}
