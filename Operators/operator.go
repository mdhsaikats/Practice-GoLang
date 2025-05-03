package main

import (
	"fmt"
)

func main() {
	var a = 15 + 14
	fmt.Println(a)

	b := 20 - 10
	fmt.Println(b)

	c := 10 * 10
	fmt.Println(c)

	d := 10 / 10
	fmt.Println(d)

	e := 2 % 2
	fmt.Println(e)

	inc := 1
	inc++
	fmt.Println(inc)

	dec := 2
	dec--
	fmt.Println(dec)

	var x string = "Saikat"
	fmt.Println(x)

	y := 5
	y -= 5
	fmt.Println(y)

	m := 5
	m += 5
	fmt.Println(m)

	n := 5
	n *= 5
	fmt.Println(n)

	k := 5
	k /= 5
	fmt.Println(k)

	var o = 5
	var l = 3
	fmt.Println(o > l) // returns 1 (true) because 5 is greater than 3

}
