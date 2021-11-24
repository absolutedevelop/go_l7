package main 

import "fmt"


func main(){

	var age int = 23
	var age_pointer *int = &age 

	fmt.Println("Age: ", age)
	fmt.Println("Age pointer: ", age_pointer)

	fmt.Println("Age from pointer: ", *age_pointer)

	*age_pointer = 24

	fmt.Println("Age from pointer: ", age)


}