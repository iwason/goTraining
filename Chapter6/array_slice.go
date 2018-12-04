package main

import "fmt"

func main() {
	values := [...]int{0, 1, 2, 3, 4}
	fmt.Println(values)
	//スライスは参照型なので、渡した先で変更すると反映される
	double(values[:])

	fmt.Println(values)
}

func double(values []int) {
	for i := 0; i < len(values); i++ {
		values[i] *= 2
	}
}
