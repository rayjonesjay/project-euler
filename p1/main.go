package main

import "fmt"

var limit int = 1_000

func main(){
	sum := 0
	for i := 3; i < limit; i++{
		if i % 3  == 0 || i % 5 == 0 {
			sum+=i
		}
	}
	fmt.Println(sum)
}
