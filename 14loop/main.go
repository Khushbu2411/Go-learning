package main

import "fmt"

func main()  {
	fmt.Println("Loop")
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thur", "Fri", "Sat"}
	fmt.Println(days)

	// for i:=0; i <len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }
	for _, day := range days {
		fmt.Printf("Day - %v\n", day)
	}

	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 8 {
			goto lco
		}
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println(rogueValue)
		rogueValue++
	}
	
	lco:
		fmt.Println("Jumping to end")
}