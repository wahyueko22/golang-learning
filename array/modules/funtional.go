package helper

import "fmt"

func CallFun(text string) string{
	fmt.Println(text + " okkk")
	return text
}

func GetValue() string{
    return "Hello from this another package"
}

func Recursive(start int, end int){
	if(start < end){
		start ++
		fmt.Print(start )
		Recursive(start, end)
	}
}

func ForLooping() {
	sum := 0
	for i:=0; i<10; i++{
		sum +=i;
	}
	fmt.Println( sum)
	sum1 :=0
	for sum1<10{
		sum1 ++
	}
	fmt.Println(sum1)
	
}