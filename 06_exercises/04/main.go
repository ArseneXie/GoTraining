package main

import "fmt"

func main() {
	var MyName string
	fmt.Println("Whar is your name? ")
	_,  err := fmt.Scan(&MyName)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Hello,",MyName)
}
