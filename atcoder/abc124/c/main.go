package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, 0)
	scanner.Buffer(buf, 10000000000)
	scanner.Split(bufio.ScanWords)
}

func ScanText() string {
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func main() {
	s := []byte(ScanText())

	cnt1 := 0
	cnt2 := 0

	for i, c := range s {
		switch i % 2 {
		case 0:
			if c == '0' {
				cnt2++
			} else {
				cnt1++
			}
		case 1:
			if c == '1' {
				cnt2++
			} else {
				cnt1++
			}
		}
	}
	if cnt1 > cnt2 {
		fmt.Println(cnt2)
	} else {
		fmt.Println(cnt1)
	}
}
