package main

import (
	"fmt"
	"log"
)

func main() {
	//ask user for number
	var number int
	for {
		fmt.Println("Please key in a number")
		_, _ = fmt.Scanln(&number)

		checkEvenOrOdd(number)
		fmt.Println("------------")

	}
}

func checkEvenOrOdd(number int) {

	if number%2 == 0 {
		fmt.Println(number, "Is an even number")
	} else if number%2 == 1 {
		fmt.Println(number, "Is an odd number")
	} else {
		log.Fatal(number%2, "unknown")
	}
}
