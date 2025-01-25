package main

import "fmt"

var limit int = 2_000_000 // two million

func main() {
	var sum int = 0
	primes := _primes(limit)
	for _, p := range primes {
		sum += p
	}
	fmt.Println(sum)
}

func _primes(lim int) []int {
	var res []int

	isPrime := make([]bool, lim+1)

	// assume all nums are primes
	for i := 2; i <= lim; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= lim; i++ {
		if isPrime[i] {
			for j := i * i; j <= lim; j += i {
				isPrime[j] = false
			}
		}
	}

	for i := 2; i <= lim; i++ {
		if isPrime[i] {
			res = append(res, i)
		}
	}

	return res
}
