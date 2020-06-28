package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, 0)
	scanner.Buffer(buf, 1000000009)
	scanner.Split(bufio.ScanWords)
}

func main() {
	scanner.Scan()
	bs := []byte(scanner.Text())

	c0 := 0
	c1 := 0
	for _, b := range bs {
		if b == '0' {
			c0++
		} else if b == '1' {
			c1++
		}
	}
	if c0 > c1 {
		fmt.Println(c1 * 2)
	} else {
		fmt.Println(c0 * 2)
	}
}
