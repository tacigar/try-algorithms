package main

import (
	"fmt"
)

var strs = []string{"dream", "dreamer", "erase", "eraser"}

func fn(s string) bool {
	slen := len(s)
	if slen == 0 {
		return true
	}
	for _, str := range strs {
		strlen := len(str)
		if slen >= strlen && s[:strlen] == str && fn(s[strlen:]) {
			return true
		}
	}
	return false
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	if fn(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
