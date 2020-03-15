package main

import (
	"fmt"
	"math"
)

type p struct {
	X, Y float64
}

func main() {
	n := 0
	fmt.Scan(&n)
	m := make([]p, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m[i].X)
	}

	if n > 1 {
		m[1].Y = math.Abs(m[1].X - m[0].X)
	}
	for i := 2; i < n; i++ {
		m[i].Y = math.Min(math.Abs(m[i].X-m[i-1].X)+m[i-1].Y, 3*math.Abs(m[i].X-m[i-2].X)+m[i-2].Y)
	}
	fmt.Print(int(m[n-1].Y))
}
