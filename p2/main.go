package main

import "fmt"

var limit int = 4_000_000

func main(){
		fmt.Println(fib(1,1))
}

func fib(a,b int) int {
	if b >= limit {
			return 0
	}
	if b % 2 == 0 {
			return b + fib(b,a+b)
	}
	return fib(b,b+a)
}
