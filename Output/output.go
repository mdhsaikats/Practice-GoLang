package main

import (
	"fmt"
)

func main() {
	a := "Saikat"
	b := "Sikder"
	c := "Goat"

	fmt.Print(a)
	fmt.Println(b)
	fmt.Printf("%v", c)

	fmt.Print(a, "Hey", b)

	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
}
