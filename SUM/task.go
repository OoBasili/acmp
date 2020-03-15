package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)
	b := a * a
	if a < 1 {
		b = 2 - b
	}
	fmt.Print((b + a) / 2)
}
