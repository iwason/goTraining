package main

import "fmt"

func main() {
	//これは配列
	values := []int{0, 1, 2, 3, 4}

	//これはスライス
	s1 := values[1:3]
	fmt.Println(s1)
	fmt.Println("len:", len(s1))
	fmt.Println("cap:", cap(s1))

	s2 := s1[1:4]
	fmt.Println(s2)
	fmt.Println("len:", len(s2))
	fmt.Println("cap:", cap(s2))

	s3 := s1[1:3:4]
	fmt.Println(s3)
	fmt.Println("len:", len(s3))
	fmt.Println("cap:", cap(s3))

}
