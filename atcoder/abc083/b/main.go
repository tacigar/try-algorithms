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
	value, _ := strconv.Atoi(scanner.Text())
	return value
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	a := scanInt()
	b := scanInt()
	sum := 0
	for v := 1; v <= n; v++ {
		val := v
		dsum := 0
		for {
			dsum += val % 10
			if val/10 == 0 {
				break
			}
			val /= 10
		}
		if dsum >= a && dsum <= b {
			sum += v
		}
	}
	fmt.Println(sum)
}
