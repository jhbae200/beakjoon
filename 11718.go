package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := ""
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		a += s.Text() + "\n"
	}
	fmt.Print(a)
}
