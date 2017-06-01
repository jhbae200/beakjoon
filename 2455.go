package main

import "fmt"

func main() {
	var total, max int
	for i := 0; i < 4; i++ {
		var a, b int
		fmt.Scanf("%d %d\n", &a, &b)
		total += b - a
		if max < total {
			max = total
		}
	}
	fmt.Println(max)
}
