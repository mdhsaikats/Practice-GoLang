package main

import (
	"fmt"
)

func main() {
	var Positon string = "Lecturer"
	var cgpa float32 = 3.1214
	var pf bool = true
	var VariableName int = 10

	var Name = "Saikat"
	var name = 30
	Variablename := 20
	pos := 200

	fmt.Println(VariableName)
	fmt.Println(name)
	fmt.Println(Variablename)
	fmt.Println(Positon)
	fmt.Println(pos)
	fmt.Println(Name)
	fmt.Println(pf)
	fmt.Println(cgpa)

	var a, b, c, d, e int = 1, 2, 3, 4, 5

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	x, y, z := 1, 2, "Saikat"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//declare in a block

	var (
		p int
		m int    = 1
		j string = "hello"
	)

	fmt.Println(p)
	fmt.Println(m)
	fmt.Println(j)

}
