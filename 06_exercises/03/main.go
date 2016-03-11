package main

import "fmt"

func main() {
	var myname string
	fmt.Print("Whar is your name? ")
	fmt.Scan(&myname)

	fmt.Println("Hello,", myname)
}
