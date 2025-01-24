package main

import "fmt"

func main() {
	var num int = 600851475143
	fmt.Println("prime factors are",findPrimeFactors(num))
}

func findPrimeFactors(n int) []int {
	var res []int
	for i := 2; i * i <= n; i++{
		// while the current number can still be divided by n, keep on dividing
		// and add the number to res
		for n % i == 0 {
			res = append(res,i)
			n = n / i
		}
	}
	if n > 1 {
		res = append(res,n)
	}
	return res
}
