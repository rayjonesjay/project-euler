package main

import (
		"fmt"
		"strconv"
)
func main(){
	var largestPalindromeNumber = 0
	// it will be efficient if we start from the largest 3 digit number going downwards
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j--{
			var res = i * j
			var strRes = strconv.Itoa(res)
			if isPalindrome(strRes) && res > largestPalindromeNumber {
					largestPalindromeNumber = res
			}
		}
	}
	fmt.Println(largestPalindromeNumber)
}

func isPalindrome(s string) bool {
		for i , j := 0, len(s)-1; i < len(s)/2; i , j = i + 1, j - 1 {
				if s[i] != s[j] {
						return false
				}
		}
		return true
}
