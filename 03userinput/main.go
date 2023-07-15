package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza")

	// comma ok // err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if ( err != nil){
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ",numrating + 1)
	}

}
