package main

import (
	"fmt"
)

func main() {

	//type string
	var username string = "Khushbu"
	fmt.Println("Variables - username: ", username)
	fmt.Printf("username is of type: %T \n", username)

	//type bool
	var isLoggedIn bool = true
	fmt.Println("Variables - isLoggedIn: ", isLoggedIn)
	fmt.Printf("isLoggedIn is of type: %T \n", isLoggedIn)

	//type int
	var smallVal int = 956435645788766789
	fmt.Println("Variables - smallVal: ", smallVal)
	fmt.Printf("smallVal is of type: %T \n", smallVal)

	//type float32
	var smallFloat float32 = 255.3453645467
	fmt.Println("Variables - smallFloat: ", smallFloat)
	fmt.Printf("smallFloat is of type: %T \n", smallFloat)

	//type float64
	var bigFloat float64 = 255.3453645467
	fmt.Println("Variables - bigFloat: ", bigFloat)
	fmt.Printf("bigFloat is of type: %T \n", bigFloat)

	// declaring type
	var anotherVariable string
	fmt.Println("Variables - anotherVariable: ", anotherVariable)
	fmt.Printf("anotherVariable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "learncode.fr" 
	fmt.Println("Variables - website: ", website)
	fmt.Printf("website is of type: %T \n", website)

	// no var style - using short variable declaration
	numberOfUsers := 30045.00
	fmt.Println("Variables - numberOfUsers: ", numberOfUsers)
	fmt.Printf("numberOfUsers is of type: %T \n", numberOfUsers)

}
