package main

import (
	"fmt"
	"sort"
)

func main() {
	b := make([]int, 9)
	for i := 0; i < 9; i++ {
		fmt.Scan(&b[i])
	}
	sort.Ints(b)
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 9; j++ {
			sum := 0
			for a := 0; a < 9; a++ {
				if a == i || a == j {
					continue
				}
				sum += b[a]
			}
			if sum == 100 {
				for a := 0; a < 9; a++ {
					if a == i || a == j {
						continue
					}
					fmt.Println(b[a])
				}
				return
			}
		}
	}
}
