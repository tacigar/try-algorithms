package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanText() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	scanner.Scan()
	a, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	return a
}

func main() {
	buf := make([]byte, 1e8)
	scanner.Buffer(buf, 1e8)
	scanner.Split(bufio.ScanWords)

	l := scanInt() - 1
	r := scanInt() - 1
	s := scanText()
	rs := []rune(s)

	var sb strings.Builder

	for i := 0; i < len(rs); i++ {
		if i >= l && i <= r {
			sb.WriteRune(rs[r-(i-l)])
		} else {
			sb.WriteRune(rs[i])
		}
	}
	fmt.Println(sb.String())
}
