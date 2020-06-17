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

	var s, t string
	if scanner.Scan() {
		s = scanner.Text()
	}
	if scanner.Scan() {
		t = scanner.Text()
	}

	c := 1
	m1 := map[byte]int{}
	m2 := map[byte]int{}

	for i := 0; i < len(s); i++ {
		if m1[s[i]] == 0 {
			if m2[t[i]] == 0 {
				m1[s[i]] = c
				m2[t[i]] = c
				c++
			} else {
				fmt.Println("No")
				return
			}
		} else {
			if m1[s[i]] != m2[t[i]] {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
