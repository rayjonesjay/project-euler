package main

import "fmt"

var limit int = 10001

func main() {
	fmt.Println(genPrimes())
}

// generate prime numbers starting from 2
func genPrimes() int {
	count := 1
	for i := 2; ; i++ {
		// check if i is prime
		if isPrime(i) {
			if count == limit {
				return i
			}
			count++
		}
	}
}

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}

	// since 2 is the only even prime number
	if n%2 == 0 {
		return false
	}

	if n % 3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
