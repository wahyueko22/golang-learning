package main

import "fmt"

func main() {
	fmt.Printf("ok")

	//if
	if true {
		fmt.Println("boolean true ")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("looping %d \n", i)
	}

	var counter int8 = 0
	for counter < 3 {
		fmt.Printf("looping %d  \n", counter)
		counter++
	}
	ArithmaticOper()
	AssignedOperartion()
	ComparationOper()
	LogicalOperator()
	BitwiseOper()
}

func ArithmaticOper() {
	var x, y = 35, 7

	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x mod y = %d\n", x%y)

	x++
	fmt.Printf("x++ = %d\n", x)

	y--
	fmt.Printf("y-- = %d\n", y)
}

func AssignedOperartion() {
	var x, y = 15, 25
	x = y
	fmt.Println("= ", x)

	x = 15
	x += y
	fmt.Printf("+=", x)

	x = 50
	x -= y
	fmt.Println("-=", x)

	x = 2
	x *= y
	fmt.Println("*=", x)

	x = 100
	x /= y
	fmt.Println("/=", x)

	x = 40
	x %= y
	fmt.Println("%=", x)
}

func ComparationOper() {
	var x, y = 15, 25

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)
}

func LogicalOperator() {
	var x, y, z = 10, 20, 30

	fmt.Println(x < y && x > z)
	fmt.Println(x < y || x > z)
	fmt.Println(!(x == y && x > z))
}

func BitwiseOper() {
	var x uint = 9  //0000 1001
	var y uint = 65 //0100 0001
	var z uint

	z = x & y
	fmt.Println("x & y  =", z)

	z = x | y
	fmt.Println("x | y  =", z)

	z = x ^ y
	fmt.Println("x ^ y  =", z)

	z = x << 1
	fmt.Println("x << 1 =", z)
	//
	z = x >> 1
	fmt.Println("x >> 1 =", z)
}
