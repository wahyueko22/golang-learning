package main

import (
	"fmt"

	"github.com/wahyueko22/golang-learning/structure/entity"
)

type rectangle struct{
	length int
	widht int
	color string
}

func main(){
	//instanciate directly
	var rect1 = rectangle{10,20,"red"}
	fmt.Println(rect1)

	var rect2 rectangle
	rect2.length=10;
	rect2.color ="blue"
	fmt.Println(rect2)

	rect3 := new(rectangle)
	rect3.length=10
	rect3.color="new blue"
	fmt.Println(rect3)

	rect4 := &rectangle{}
	rect4.widht=5
	rect4.color="yellow pointer"
	fmt.Println(rect4)

	rect5 := &rectangle{}
	(*rect5).widht=5
	(*rect5).color="red pointer"
	fmt.Println(rect5)

	rect6:=&rect2
	rect6.color = "color_6"
	fmt.Println(rect6, rect2)

	emp1 := entity.Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []entity.Salary{
			entity.Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			entity.Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			entity.Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	emp1.EmpInfo()
	
}