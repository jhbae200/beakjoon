package main

import ."fmt"

func main() {
	var x, c int
	Scan(&x)
	for i := 64; i > 0; i/=2 {
		if i <= x {
			c++
			x -= i
		}
	}
	Print(c)
}
