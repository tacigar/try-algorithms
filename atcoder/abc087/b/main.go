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
	a := scanInt()
	b := scanInt()
	c := scanInt()
	x := scanInt()
	cnt := 0
	for ai := 0; ai <= a; ai++ {
		for bi := 0; bi <= b; bi++ {
			for ci := 0; ci <= c; ci++ {
				if ai*500+bi*100+ci*50 == x {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
