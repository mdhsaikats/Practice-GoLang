package main

import (
	"fmt"
)

func main() {
	var a = map[string]int{"Saikat": 22, "Sikder": 23}
	fmt.Println(a["Saikat"])
	fmt.Println(a["Sikder"])

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("a\t%T\n", a)
	fmt.Printf("a\t%#v\n", a)

	b := map[string]int{"Saikat": 22, "Sikder": 23}
	fmt.Println(b["Saikat"])
	fmt.Println(b["Sikder"])

	var c = make(map[string]string)
	var d map[string]string

	fmt.Println(c == nil) // false
	fmt.Println(d == nil) // true
}
