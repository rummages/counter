package main 

import "fmt"

func main() { 
	i := 1 
	for i <= 100 { // For loop creates another little function inside main
		fmt.Println(i)
		i = i + 1 // Here we have to initialize the method to increment the loop
		} // Closes for loop
	} // Closes function

