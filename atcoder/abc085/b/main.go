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
	m := map[int]struct{}{}
	for i := 0; i < n; i++ {
		m[scanInt()] = struct{}{}
	}
	fmt.Println(len(m))
}
