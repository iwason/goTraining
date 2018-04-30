package main

import "fmt"
import "unicode/utf8"

func main() {
	ja := "Go言語"
	fmt.Println(ja, "len:", utf8.RuneCountInString(ja))
	fmt.Println(ja, "len:", utf8.RuneCountInString(ja))
}
