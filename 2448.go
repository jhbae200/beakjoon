package main

import (
	. "fmt"
	"strings"
)

func z(c [][]string, n int, x int, y int) {
	if n == 3 {
		c[x][y] = "*"
		c[x+1][y] = "*"
		c[x+1][y+2] = "*"
		c[x+2][y] = "*"
		c[x+2][y+1] = "*"
		c[x+2][y+2] = "*"
		c[x+2][y+3] = "*"
		c[x+2][y+4] = "*"
	} else {
		z(c, n/2, x, y)
		z(c, n/2, x+n/2, y)
		z(c, n/2, x+n/2, y+n)
	}

}

func main() {
	var n int
	Scan(&n)
	ch := make([][]string, n)
	for i := 0; i < len(ch); i++ {
		ch[i] = make([]string, n+i)
		for j := 0; j < n+i; j++ {
			ch[i][j] = " "
		}
	}
	z(ch, n, 0, 0)
	for i := 0; i < len(ch); i++ {
		for j := i; j < n-1; j++ {
			print(" ")
		}
		println(strings.Join(ch[i],""))
	}
}
