package main

import "fmt"

func main() {
	fmt.Println(pythag())
}

func pythag() (res int) {
	for a := 1; a <= 1000; a++{
		for b := a + 1 ; b <= 1000; b++{
			c := 1000 - a - b 
			if a*a + b*b == c*c {
				res = a * b * c
				break
			}
		}
	}
	return res
}
