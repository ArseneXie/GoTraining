package main

import "fmt"

func main() {
	fmt.Println("Print all the even number form 0 to 100")
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
		continue
	}
}
