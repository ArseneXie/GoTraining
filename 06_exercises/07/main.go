package main

import "fmt"

func main() {
	fmt.Println("Sum all the mul of 3 and 5 under 1000")

	sum := 0

	for i := 1 ; i < 1000; i++ {

		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
			// or sum += i
			//fmt.Println(i)
		}
	}

	fmt.Println("Sum of 3,5 multiple as ", sum)
}
