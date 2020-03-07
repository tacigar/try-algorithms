package main

import (
	"fmt"
	"math"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var s string
	fmt.Scan(&s)

	space := make([]byte, 5*int(math.Pow(10, 5)))
	si := 2 * int(math.Pow(10, 5))
	ei := si + len(s)
	for i := si; i < ei; i++ {
		space[i] = s[i-si]
	}

	b := false

	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var t int
		fmt.Scan(&t)
		if t == 1 {
			b = !b
		} else { // t == 2
			var f int
			var c string
			fmt.Scan(&f, &c)
			if f == 1 && !b || f == 2 && b { // 先頭
				si--
				space[si] = c[0]
			} else {
				space[ei] = c[0]
				ei++
			}
		}
	}

	str := string(space[si:ei])
	if !b {
		fmt.Println(str)
	} else {
		fmt.Println(reverse(str))
	}
}
