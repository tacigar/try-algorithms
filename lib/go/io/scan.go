package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
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
