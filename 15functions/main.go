package main

import "fmt"

func main()  {
	greet()
	result := adder(3,5)
	fmt.Println("Functions")
	fmt.Println(result)

	resultOne, message := proAdded(3,5,4,5,2,1,34)
	fmt.Println(resultOne, message)
}

func greet()  {
	fmt.Println("Namaste")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdded(value ...int) (int, string) {
	total := 0
	for _,i := range(value) {
		total += i
	}
	return total, "HI"
}