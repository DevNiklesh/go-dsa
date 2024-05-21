package main

import (
	"fmt"
	"strconv"
)

func checkPalindrome(n int) string {
	yes := "Palindrome number"
	no := "Not a palindrome"

	str := strconv.Itoa(n)
	for i := range str {
		if str[i] != str[len(str)-i-1] {
			return no
		}
	}

	return yes
}

func checkPali(n int) string {
	yes := "Palindrome number"
	no := "Not a palindrome"

	revNum := 0
	dup := n

	for n > 0 {
		revNum = revNum*10 + n%10
		n = n / 10
	}

	if revNum == dup {
		return yes
	}

	return no
}

func main() {
	fmt.Println(checkPalindrome(4554))
	fmt.Println(checkPali(7798))
	fmt.Println(checkPali(4554))
}
