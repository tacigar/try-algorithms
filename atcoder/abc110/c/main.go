package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	scanner.Buffer(buf, 1000000009)
	scanner.Split(bufio.ScanWords)

	var s, t []rune
	if scanner.Scan() {
		s = []rune(scanner.Text())
	}
	if scanner.Scan() {
		t = []rune(scanner.Text())
	}

	ss := map[rune]rune{}
	tt := map[rune]rune{}

	for i := 0; i < len(s); i++ {
		if _, ok := ss[s[i]]; !ok {
			ss[s[i]] = t[i]
		}
		if _, ok := tt[t[i]]; !ok {
			tt[t[i]] = s[i]
		}
		if ss[s[i]] != t[i] || tt[t[i]] != s[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
