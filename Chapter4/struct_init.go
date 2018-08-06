package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Employee struct {
	id int
	Person
}

func main() {
	var p1 Person
	p1.name = "John"
	p1.age = 10

	p2 := Person{name: "moge", age: 12}
	p3 := Person{"hoge", 111}
	//構造体ポインタ型？
	p4 := &Person{"Mike", 30}
	fmt.Println(p1, p2, p3, p4)
	fmt.Println(p3.name)
	fmt.Println(p4.name)

	p3.name = "copy man"
	p4.name = "copy man"

	fmt.Println(p3, p4)

	e := Employee{1, Person{"takeshi", 20}}
	fmt.Println(e)

}
