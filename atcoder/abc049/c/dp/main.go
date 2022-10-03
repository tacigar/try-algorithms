package main

import (
	"fmt"
)

var strs = []string{"dream", "dreamer", "erase", "eraser"}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	slen := len(s)

	var dp [200000]bool
	dp[0] = true
	for i := 0; i < slen; i++ {
		if !dp[i] {
			continue
		}
		for _, str := range strs {
			strlen := len(str)
			if i+strlen <= slen && str == s[i:i+strlen] {
				dp[i+strlen] = true
			}
		}
	}

	if dp[slen] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
