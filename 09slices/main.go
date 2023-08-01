package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Slices")
	var fruitList = []string{"Peach", "Mango", "Apple"}
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Tomato", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 456
	highScores[2] = 345
	highScores[3] = 486
	//highScores[4] = 215
	highScores = append(highScores, 555, 768, 123)
	
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	courses := []int{1,2,3,4,5}
	index := 2 
	courses = append(courses[:index], courses[index+1:]... )

	fmt.Println(courses)



}