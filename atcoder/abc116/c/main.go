package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	scanner.Split(bufio.ScanWords)
}

func ScanInt() int {
	scanner.Scan()
	t := scanner.Text()
	n, _ := strconv.Atoi(t)
	return n
}

func main() {
	n := ScanInt()

	hs := make([]int, n)
	for i := 0; i < n; i++ {
		hs[i] = ScanInt()
	}

	maxH := 0
	for _, h := range hs {
		if maxH < h {
			maxH = h
		}
	}

	cnt := 0
	for h := 1; h <= maxH; h++ {
		b := false
		for i := 0; i < n; i++ {
			if hs[i] >= h && !b {
				cnt++
				b = true
			} else if hs[i] < h {
				b = false
			}

		}
	}
	fmt.Println(cnt)
}
