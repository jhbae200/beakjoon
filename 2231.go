package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	for i := int(float64(a) * 0.8); i < a; i++ {
		sum := i
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		if sum == a {
			fmt.Print(i)
			return
		}
	}
	fmt.Print(0)
}
