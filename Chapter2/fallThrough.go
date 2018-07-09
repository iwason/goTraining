package main

import (
	"fmt"
	"time"
)

func fallThrough() {
	for day := time.Sunday; day <= time.Saturday; day++ {
		switch day {

		case time.Sunday:
			fallthrough
		case time.Saturday:
			fmt.Println(day, "休日")
		default:
			fmt.Println(day, "平日")
		}
	}
}
