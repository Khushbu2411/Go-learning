package main

import "fmt"

func main()  {
	fmt.Println("Array")

	var fruitList [4]string
	fruitList[0] = "Mango"
	fruitList[1] = "Watermelon"
	fruitList[3] = "Apple"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Mushroom", "Spinach"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))

}