package main

import "fmt"


type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func main()  {
	fmt.Println("Structs....")
	khushbu := User{"Khushbu", "khushbuag2411@gmail.com", true, 24}
	fmt.Println(khushbu)
	fmt.Printf("khushbu's details are: %+v\n", khushbu)
	fmt.Printf("Name is %v and email is %v.\n", khushbu.Name, khushbu.Email)

	
}