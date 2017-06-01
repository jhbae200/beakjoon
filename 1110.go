package main

import "fmt"

func main() {
	var a, d int
	fmt.Scan(&a)
	if a < 10 {
		a *= 10
	}
	c := a
	for {
		b := c/10 + c%10
		c = c%10*10 + b%10
		d++
		if c == a {
			break
		}
	}
	fmt.Print(d)
}
