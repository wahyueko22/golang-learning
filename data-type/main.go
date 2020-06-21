package main

import (
	"fmt"
	"reflect"
)

func main(){
	//integer
	fmt.Println("ok")
	var intDefault int
	fmt.Printf("default of int : %d \n", intDefault)
	var a int8   = 9
	fmt.Printf("value of variable a : %d \n", a)
	fmt.Printf(" 'a' varible type %s\n", reflect.TypeOf(a))
	var intA int8 = 'a'
	fmt.Printf("intA = %d\n",intA)

	//alias
	var char byte = 'a' 
	fmt.Printf("char = %d\n",char)
	fmt.Printf("Character: %c\n", char)
	var runeTest rune = '♥'
	fmt.Printf("data type rune : %c\n", runeTest)

	//floating
	var floatingOne float32 = 10.4
	fmt.Printf("floating example : %.2f \n", floatingOne)

	//string
	str := "hallo"
	fmt.Printf("string : %s \n", str)

	//boolean
	isBol := true
	if(isBol){
		fmt.Printf("boolean true")
	}

}