package main

import "fmt"

func main()  {
	fmt.Println("Pointers....")

	var ptr *int
	fmt.Println("Value of pointer is:", ptr)

	myNumber := 23
	var ptrv = &myNumber
	fmt.Println("pointer ptrv is: ", ptrv)	
	fmt.Println("value of pointer ptrv is: ", *ptrv)	

	*ptrv = *ptrv * 2 
	fmt.Println("New value of pointer ptrv is: ", *ptrv)	

}