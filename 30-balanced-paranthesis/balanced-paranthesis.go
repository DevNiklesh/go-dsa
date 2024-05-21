package main

import "fmt"

func checkBalancedParanthesis(s string) bool {
	if len(s) < 2 {
		return false
	}

	openBrackets := []string{}
	bracketMap := map[string]string{")": "(", "}": "{", "]": "["}

	for i := 0; i < len(s); i++ {
		char := string(s[i])
		switch char {
		case "(", "{", "[":
			openBrackets = append(openBrackets, char)
		case ")", "}", "]":
			length := len(openBrackets)
			if length == 0 || openBrackets[length-1] != bracketMap[char] {
				return false
			}
			openBrackets = openBrackets[:length-1]
		}
	}

	return len(openBrackets) == 0
}

func main() {
	fmt.Println(checkBalancedParanthesis("()[{}()]"))
	fmt.Println(checkBalancedParanthesis("[)(]"))
	fmt.Println(checkBalancedParanthesis("["))
	fmt.Println(checkBalancedParanthesis("(())"))
	fmt.Println(checkBalancedParanthesis("(("))
}
