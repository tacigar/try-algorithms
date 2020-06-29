package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	_, _ = fmt.Fprintf(os.Stderr, "Text not found")
	return ""
}

func ScanInt() (ret int) {
	var err error
	if scanner.Scan() {
		txt := scanner.Text()
		if ret, err = strconv.Atoi(txt); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, err.Error())
		}
	}
	return
}

func main() {
	n, q := ScanInt(), ScanInt()
	s := ScanText()

	cur := 0
	t := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i-1] == 'A' && s[i] == 'C' {
			cur++
		}
		t[i] = cur
	}

	var w = bufio.NewWriter(os.Stdout)
	for i := 0; i < q; i++ {
		l, r := ScanInt()-1, ScanInt()-1
		_, _ = fmt.Fprintln(w, t[r]-t[l])
	}
	_ = w.Flush()
}
