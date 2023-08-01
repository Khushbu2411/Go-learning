package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("welcome")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // for formatting important to use this time date and day

	createdDate := time.Date(1997, time.January, 24, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))


}