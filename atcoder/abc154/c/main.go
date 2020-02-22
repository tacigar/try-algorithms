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

	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[scanInt()]++
	}

	for _, v := range m {
		if v > 1 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
