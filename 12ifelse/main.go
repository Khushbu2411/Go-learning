package main

import "fmt"


func main()  {
	fmt.Println("Go If else")
	loginCount := 10
	var result string
	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	}else {
		result = "Exactly 10"
	}
	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	}

	//if err != nil {

	//}
}