package main

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

func ScanInt() int {
	scanner.Scan()
	txt := scanner.Text()
	n, _ := strconv.Atoi(txt)
	return n
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Info struct{ X, Y, H int }

func main() {
	n := ScanInt()
	infoList := make([]Info, n)
	for i := 0; i < n; i++ {
		infoList[i].X = ScanInt()
		infoList[i].Y = ScanInt()
		infoList[i].H = ScanInt()
	}

	for cx := 0; cx <= 100; cx++ {
		for cy := 0; cy <= 100; cy++ {
			var h int
			for _, info := range infoList {
				if info.H != 0 {
					h = info.H + IntAbs(info.X-cx) + IntAbs(info.Y-cy)
					break
				}
			}
			success := true
			for _, info := range infoList {
				if info.H != IntMax(h-IntAbs(info.X-cx)-IntAbs(info.Y-cy), 0) {
					success = false
					break
				}
			}
			if success {
				fmt.Printf("%d %d %d\n", cx, cy, h)
				return
			}
		}
	}
}
