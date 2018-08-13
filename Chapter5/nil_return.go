package main

import "fmt"

/* sample class */
type MyError string

func (e MyError) Error() string {
	return string(e)
}
func main() {
	err := do()

	if err == nil {
		fmt.Println("no error")
	} else {
		fmt.Println("with error??", err)
	}
}
func do() error {
	var ret *MyError = nil
	return ret
}

/*
error interfaceの定義
type error interface {
    Error() string
}
*/
