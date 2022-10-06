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
	cs := make([][]int, 3)
	for i := 0; i < 3; i++ {
		cs[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cs[i][j] = scanInt()
		}
	}
	for i := 0; i < 2; i++ {
		d := cs[0][i] - cs[0][i+1]
		for j := 1; j < 3; j++ {
			if d != cs[j][i]-cs[j][i+1] {
				fmt.Println("No")
				return
			}
		}
	}
	for i := 0; i < 2; i++ {
		d := cs[i][0] - cs[i+1][0]
		for j := 1; j < 3; j++ {
			if d != cs[i][j]-cs[i+1][j] {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
