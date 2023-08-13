package main

import "fmt"

func main()  {
	defer fmt.Println("hello world")
	defer fmt.Println("hellosef world")
	fmt.Println("hello")
	fmt.Println("2hello")
	fmt.Println("h1ello")
	myDefer()
}
func myDefer()  {
	for i := 0 ; i < 5; i ++ {
		defer fmt.Println(i)
	}
	
}