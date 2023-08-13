package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Switch case")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is:", diceNumber)
	switch diceNumber {
		case 1:
			fmt.Println("Open")
		case 2:
			fmt.Println("Move 2")
		case 3:
			fmt.Println("Move 3")
			fallthrough //execute next case also. get case 3 and case 4 is also executed.
		case 4:
			fmt.Println("Move 4")
		case 5:
			fmt.Println("Move 5")
		case 6:
			fmt.Println("Move and roll agian - 6") 

	}

}