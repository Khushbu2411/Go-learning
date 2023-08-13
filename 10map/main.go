package main

import (
	"fmt"
)

func main()  {
	fmt.Println("MAPssss")
	languages := make(map[string]string)
	languages["RB"] = "Ruby"
	languages["PY"]  = "Python"
	languages["JS"]  = "JavScript"
	fmt.Println(languages)
	fmt.Println(languages["RB"])

	//loops 
	for key, value := range languages {
		fmt.Printf("For key %v - value %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("value  %v\n", value)
	}

	delete(languages, "PY")
	fmt.Println(languages)

	

	
}