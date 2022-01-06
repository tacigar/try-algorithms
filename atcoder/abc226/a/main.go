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

func main() {
	scanner.Split(bufio.ScanWords)
	x := scanText()
	ss := strings.Split(x, ".")
	n := ss[1][0] - '0'
	d, _ := strconv.Atoi(ss[0])
	if n >= 0 && n <= 4 {
		fmt.Println(d)
	} else {
		fmt.Println(d + 1)
	}
}
