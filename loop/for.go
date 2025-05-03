package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Saikat")
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	//Nested Loops

	size := [2]string{"BIG", "SMALL"}
	fruits := [3]string{"Apple", "Orange", "Banana"}
	for i := 0; i < len(size); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(size[i], ":", fruits[j])
		}
	}

	//range
	name := [3]string{"Saikat", "Sikder", "DIU"}
	for index, v := range name {
		fmt.Printf("%v\t%v\n", index, v)
	}

	//omit the index
	id := [2]int{449, 660}
	for _, v := range id {
		fmt.Printf("%v\n", v)
	}

	//omit the values
	value := [2]int{1, 2}
	for idx, _ := range value {
		fmt.Printf("%v\n", idx)
	}
}
