package main

import "fmt"

type myType int

func main() {
	var z myType = 123

	fmt.Println(z.add(3))
	f := z.add
	fmt.Println(f(3))
	//足されたままの数値
	fmt.Println(z.add(0))

}

func (value *myType) add(increment myType) myType {
	*value += increment
	return *value
}
