package persegi

import "fmt"

//exported function
func Keliling(sisi int) {
	keliling := sisi * 4
	fmt.Printf("sisi persegi : %d \n", keliling)
	printEnd()
}

//unexported function
func printEnd() {
	fmt.Printf(" === end === \n")
}
