package main

import "fmt"

func main() {
	fmt.Println("Print 3 mul as Fuzz, 5 mul as Buzz  from 1 to 100")

	//astring :=""
	//bstring:=""

	var astring string
	var bstring string

	for i := 1; i <= 100; i++ {

		if i%3 == 0 {
			astring = "Fizz"
		} else {
			astring = ""
		}
		if i%5 == 0 {
			bstring = "Buzz"
		} else {
			bstring = ""
		}
		fmt.Println(i, " show as ", astring+bstring)

	}
}
