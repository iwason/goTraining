package main

import "fmt"

func length() {
	en := "golang"
	ja := "Go言語"

	fmt.Println(en, " len: ", len(en))
	fmt.Println(ja, " len: ", len(ja))
}
