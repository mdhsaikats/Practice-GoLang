package main

import (
	"fmt"
)

func myMassage() {
	fmt.Println("Hello")
}

func sum(y int) {
	x := y + 10
	fmt.Println(x)
}

func name(name string, age int) {
	fmt.Println("Hello", name, age, "Lad")
}

func Myreturn(x int, y int) int {
	return x + y
}

func main() {
	myMassage()
	sum(10)
	name("Saikat", 22)
	fmt.Println(Myreturn(1, 2))
}
