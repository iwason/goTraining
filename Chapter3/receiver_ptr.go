package main

import "fmt"

type myType int

func (value myType) setByValue(newValue myType) {
	value = newValue
}
func (value *myType) setByPointer(newValue myType) {
	*value = newValue
}

func (value *myType) add(increment myType) myType {
	*value += increment
	return *value
}

func main() {
	var z myType

	z.setByValue(1)

	fmt.Println(z)

	z.setByPointer(33)
	fmt.Println(z)

	z.add(3)
	fmt.Println(z)
	//メソッド値
	f := z.add
	fmt.Println(f(33))
	fmt.Println(z)

}
