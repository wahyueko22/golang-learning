package main

import "fmt"

func main() {

	forLooping()
	BoolLooping()

}

//for looping
func forLooping() {
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
}

//while looping
func BoolLooping() {
	isBol := false
	counter := 0
	for !isBol {
		if counter > 2 {
			isBol = true

		}
		fmt.Println(counter)
		counter++
	}
}
