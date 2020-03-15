package main

import "fmt"

var a []int

func d(i int) (s int) {
	for i >= 0 {
		s += a[i]
		i = (i & (i + 1)) - 1
	}
	return
}

func main() {
	var n, k, s, v, i int
	fmt.Scan(&n, &k)
	for i < k {
		a = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&v)
			v--
			s += d(n-1) - d(v)
			for v < n {
				a[v]++
				v = v | (v + 1)
			}
		}
		i++
	}
	fmt.Print(s)
}
