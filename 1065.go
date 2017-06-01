package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b := a
	if a >= 100 {
		b = 99
	}
	for i := 111; i <= a; i++ {
		if 1000==i {
			break
		}
		c := make([]int, 3)
		k := 0
		for j := i; j > 0; j /= 10 {
			c[k] = j%10
			k++
		}
		if c[0]-c[1]==c[1]-c[2] {
			b++
		}
	}
	fmt.Print(b)
}
