package main

import "fmt"

func main() {
	var a,b int
	fmt.Scan(&a)
	b=a
	if a%5 == 0 {
		b = a/5
	} else {
		for i:=0; a>5*i; i++  {
			if (a-5*i)%3 == 0 {
				b = i+(a-5*i)/3
			}
		}
	}
	if a==b {
		b = -1
	}
	fmt.Print(b)
}
