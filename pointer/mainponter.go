package main

import "fmt"

func main() {
    b := 6 
 //& is the address of operator. * represent a pointer in some cases, in others it is used as the 'dereference operator'.
    var b_ptr *int // *int is used delcare variable
                   // b_ptr to be a pointer to an int

    b_ptr = &b     // b_ptr is assigned value that is the address
                       // of where variable b is stored

    // Shorhand for the above two lines is:
    // b_ptr := &b

    fmt.Printf("address of b_ptr: %p\n", b_ptr)

    // We can use *b_ptr get the value that is stored
    // at address b_ptr, or dereference the pointer 
	fmt.Printf("value stored at b_ptr: %d\n", *b_ptr)
	
	var a = 5
    var p = &a // copy by reference
    var x = a  // copy by value

    fmt.Println("a = ", a)   // a =  5
    fmt.Println("p = ", p)   // p =  0x10414020
    fmt.Println("*p = ", *p) // *p =  5
    fmt.Println("&p = ", &p) // &p =  0x1040c128
    fmt.Println("x = ", x)   // x =  5

    fmt.Println("\n Change *p = 3")
    *p = 3
    fmt.Println("a = ", a)   // a =  3
    fmt.Println("p = ", p)   // p =  0x10414020
    fmt.Println("*p = ", *p) // *p =  3
    fmt.Println("&p = ", &p) // &p =  0x1040c128
    fmt.Println("x = ", x)   // x =  5

    fmt.Println("\n Change a = 888")
    a = 888
    fmt.Println("a = ", a)   // a =  888
    fmt.Println("p = ", p)   // p =  0x10414020
    fmt.Println("*p = ", *p) // *p =  888
    fmt.Println("&p = ", &p) // &p =  0x1040c128
    fmt.Println("x = ", x)   // x =  5

    fmt.Println("\n Change x = 1")
    x = 1
    fmt.Println("a = ", a)   // a =  888
    fmt.Println("p = ", p)   // p =  0x10414020
    fmt.Println("*p = ", *p) // *p =  888
    fmt.Println("&p = ", &p) // &p =  0x1040c128
    fmt.Println("x = ", x)   // x =  1

    &p = 3 // error: Cannot assign to &p because this is the address of variable a

}