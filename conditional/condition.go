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

	bol := true
	if bol {
		fmt.Println("masuk2")
	}

}
