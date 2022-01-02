package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanInt() int {
	scanner.Scan()
	a, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	return a
}

func main() {
	scanner.Split(bufio.ScanWords)

	n := scanInt()
	x := scanInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		l := scanInt()
		ai := make([]int, l)
		for j := 0; j < l; j++ {
			ai[j] = scanInt()
		}
		a[i] = ai
	}

	cnt := 0
	var rec func(acc, i int)
	rec = func(acc, i int) {
		if i == len(a) {
			if acc == x {
				cnt++
			}
			return
		}
		if acc > x {
			return
		}
		ai := a[i]
		for _, v := range ai {
			rec(acc*v, i+1)
		}
	}
	rec(1, 0)
	fmt.Println(cnt)
}
