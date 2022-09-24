package main

import (
	"bufio"
	"fmt"
	"math"
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
	min := 100000000
	for i := 0; i < n; i++ {
		a := scanInt()
		cnt := 0
		for {
			if a&1 != 0 {
				break
			}
			a = a >> 1
			cnt++
		}
		min = int(math.Min(float64(min), float64(cnt)))
	}
	fmt.Println(min)
}
