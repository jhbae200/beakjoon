package main

import . "fmt"

func main() {
	var a,c int
	var s float64
	Scan(&a)
	b := make([]int, a)
	for i := 0; i < a; i++ {
		Scan(&b[i])
		if c<b[i] {
			c=b[i]
		}
	}
	for i := 0; i <a ; i++  {
		s += float64(b[i])/float64(c)*100
	}
	Printf("%.2f", s/float64(a))
}
