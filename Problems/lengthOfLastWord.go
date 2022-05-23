package main

import "fmt"

func lengthOfLastWord(s string) int {
	var isChar = false
	var length = 1
	for i := len(s) - 1; i >= 0; i-- {
		if isChar {
			if s[i] == ' ' {
				return length
			} else {
				length++
			}
		} else {
			if s[i] != ' ' {
				isChar = true
			}
		}
	}
	return length
}

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}
