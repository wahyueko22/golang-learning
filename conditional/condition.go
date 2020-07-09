package main

import "fmt"

func main() {
	IfCondition()
}

//
func IfCondition() {
	data := 1
	if data < 10 {
		fmt.Println("masuk")
	}

	//boolean condition
	bol := true
	if bol {
		fmt.Println("masuk2")
	}

	if result := 10 / 2; result < 10 {
		fmt.Printf("result : %d \n", result)
	}

}
