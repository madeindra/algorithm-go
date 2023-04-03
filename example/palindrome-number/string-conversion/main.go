package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := isPalindrome(121)
	fmt.Println(result)
}

func isPalindrome(num int) bool {
	// convert integer to string
	str := strconv.Itoa(num)

	// create array of rune/char from string (because string is immutable)
	rev := []rune(str)

	// swap character (with two pointers approach)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}

	// compare reversed string with initial string
	return string(rev) == str
}
