package main

import "fmt"

//slice declaration always with empty array []
func main() {
	//cretaing slice using make buildin function type, len, capacity
	i := make([]int, 5, 5)
	fmt.Println(i)

	fruits := []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits)
	var aFruits = fruits[0:3]
	fmt.Println(aFruits)
	var bFruits = fruits[1:4]
	fmt.Println(bFruits)
	var cfruits = append(fruits, "papaya")
	fmt.Println(cfruits)

	students := []string{"ara", "fanta", "jevan"}
	fmt.Println(students)
	//remove slice
	students = students[0:2]
	fmt.Println(students)
	students = append(students, "misa")
	fmt.Println(students)
	fmt.Println(len(students))
	for index := 0; index < len(students); index++ {
		fmt.Println("ok")
	}
	testSliceNotPointer(students)
	fmt.Println(students)
	testSlice(&students)
	fmt.Println(students)
}

//without pointer value after append not visible from outside function
func testSliceNotPointer(tempStudent []string) {
	tempStudent[0] = "zoro"
	tempStudent = append(tempStudent, "dona")

	fmt.Println(tempStudent)
}

//with pointer value after append notvisible from outside function
func testSlice(tempStudent *[]string) {
	(*tempStudent)[0] = "bams"
	(*tempStudent) = append((*tempStudent), "prince")

	fmt.Println(tempStudent)
}
