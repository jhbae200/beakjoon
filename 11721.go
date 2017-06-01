package main

import (
	."fmt"
)

func main() {
	var s string
	Scan(&s)
	for len(s) >= 10 {
		Print(s[:10]+"\n")
		s=s[10:]
	}
	Print(s)
}
