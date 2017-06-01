package main

import (
	b"bufio"
	"fmt"
	"os"
	."strings"
)

func main() {
	s := b.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 4096), 1000000)
	c := 0
	for s.Scan() {
		c+=len(Fields(s.Text()))
	}
	fmt.Print(c)

}
