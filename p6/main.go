package main

import (
	"fmt"
	"math"
)

var end int = 100

func main() {
	b := sumSq(end)
	a := sqSum(end)
	fmt.Println(a-b)
}

func sqSum(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func sumSq(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum = sum + int(math.Pow(float64(i), 2))
	}
	return
}
