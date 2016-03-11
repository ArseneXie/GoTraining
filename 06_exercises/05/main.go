package main

import "fmt"

func main() {
	var numone int
	var numtwo int

	fmt.Print("Please enter two  numbers:  ")
	fmt.Scanln(&numone,&numtwo)

	if  numone>numtwo {
		fmt.Println("Max mod min is ", numone%numtwo)
	}else{
		fmt.Println("Max mod min is ", numtwo%numone)
	}

}
