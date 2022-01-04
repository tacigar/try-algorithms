package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanText() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner.Split(bufio.ScanWords)
	p := map[string]string{
		"ox": "oxxoxxoxxoxx",
		"xx": "xxoxxoxxoxxo",
		"xo": "xoxxoxxoxxox",
	}
	s := scanText()
	if len(s) == 1 {
		fmt.Println("Yes")
	} else if _, ok := p[s[:2]]; !ok {
		fmt.Println("No")
	} else if s == (p[s[:2]])[:len(s)] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
