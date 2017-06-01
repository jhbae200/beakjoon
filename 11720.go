package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	s.Scan()
	a := 0
	for _, b := range s.Bytes() {
		a += int(b) - 48
	}
	fmt.Print(a)
}
