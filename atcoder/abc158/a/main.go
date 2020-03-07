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
	s := scanText()
	b := s[0]
	for i := 1; i < 3; i++ {
		if s[i] != b {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
