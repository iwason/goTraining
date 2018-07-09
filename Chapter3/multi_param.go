package main

import "fmt"

func main() {
	f2(f1())

	holiday(1, "元旦", "成人の日")
	holiday(2, "建国記念日")
	holiday(3, "春分の日")

}

func f1() (int, string, float32) {
	return 0, "xyz", 3.14
}

func f2(a int, b string, c interface{}) {
	fmt.Println(a, b, c)
}

func holiday(month int, days ...string) {
	fmt.Printf("%d月の祝日は%d日あります。\n", month, len(days))
	for _, day := range days {
		fmt.Println(day)
	}
	fmt.Println()

}
