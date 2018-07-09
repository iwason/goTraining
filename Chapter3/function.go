package main

func plus(a int, b int) int {
	return a + b
}
func calc(a int, b int) (int, int, int, float32) {
	return a + b, a - b, a * b, float32(a) / float32(b)
}
func main() {
	c := plus(10, 20)
	println(c)
	//ブランク演算子
	add, div, _, _ := calc(10, 4)
	println(add, div)
}
