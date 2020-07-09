package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 7
	reflectValue := reflect.ValueOf(number)

	fmt.Println("tipe  variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}
}
