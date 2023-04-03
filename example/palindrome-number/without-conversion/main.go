package main

import "fmt"

func main() {
	result := isPalindrome(121)
	fmt.Println(result)
}

func isPalindrome(num int) bool {
	// if it is 0, return false
	if num < 0 {
		return false
	}

	// 1-9 will always be palindrome, return true
	if num < 10 {
		return true
	}

	// set an initial number
	rev := 0

	// reverse number by keep dividing the number by 10
	for base := num; base > 0; base = base / 10 {
		// multiply initial number by 10 and add base
		rev = (rev * 10) + (base % 10)
		// this will add latest digit of num to rev (last digit calculated from base divided 10 modulo 10)
		// next iteration that digit will be shifted to the left (by multiplying it by 10)
	}

	// compare reversed number with num
	return rev == num
}
