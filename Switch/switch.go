package main

import (
	"fmt"
)

func main() {
	for {
		var option int
		fmt.Print("Enter your option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Hi")
		case 2:
			fmt.Println("Bye")
		case 3:
			fmt.Println("Exiting")
			return
		default:
			fmt.Println("Error!")
		}
	}
}
