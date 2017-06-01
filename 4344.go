package main

import "fmt"

func main() {
	var a,b int
	fmt.Scan(&a)
	t := make([]float64, a)
	for i := 0; i < a; i++ {
		d := 0
		fmt.Scan(&b)
		c := make([]int, b)
		for j := 0; j < b; j++ {
			fmt.Scan(&c[j])
			d+=c[j]
		}
		v := float64(d)/float64(b)
		f := 0
		for _, e := range c {
			if v < float64(e) {
				f++
			}
		}
		t[i] = float64(f)/float64(b)*100
	}
	for _, g := range t {
		fmt.Printf("%.3f%%\n", g)
	}
}
