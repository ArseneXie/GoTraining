package main

import "fmt"

func main() {
	var minnum int
	var maxnum int

	fmt.Print("Please enter a min number:  ")
	fmt.Scanln(&minnum)

	fmt.Print("Please enter a Max number:  ")
	fmt.Scanln(&maxnum)

	fmt.Println("What you enter min is ",minnum, " and the Max is ", maxnum)

	fmt.Println("Max mod min is ", maxnum%minnum)
}
