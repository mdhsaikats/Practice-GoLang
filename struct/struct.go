package main

import (
	"fmt"
)

type person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	p1 := person{
		name:   "Saikat",
		age:    22,
		job:    "Student",
		salary: 1000,
	}
	p2 := person{
		name:   "Sikder",
		age:    23,
		job:    "Student",
		salary: 2000,
	}
	fmt.Println(p1.name, p1.age, p1.job, p1.salary)
	fmt.Println(p2.name, p2.age, p2.job, p2.salary)

	person1 := person
	person2 := person

	person1.name = "Saikat"
	person1.age = 22
	person1.job = "Student"
	person1.salary = 1000

	person2.name = "Sikder"
	person2.age = 23
	person2.job = "Student"
	person2.salary = 2000

	printPerson(person1)
	printPerson(person2)
}

func printPerson(p person) {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Job:", p.job)
	fmt.Println("Salary:", p.salary)
}
