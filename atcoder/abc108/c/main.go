package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var n, k int
	if scanner.Scan() {
		tmp := scanner.Text()
		n, _ = strconv.Atoi(tmp)
	}
	if scanner.Scan() {
		tmp := scanner.Text()
		k, _ = strconv.Atoi(tmp)
	}

	if k%2 == 0 {
		var rs1, rs2 []int
		for i := 1; i <= n; i++ {
			if i%k == 0 {
				rs1 = append(rs1, i)
			} else if i%k == k/2 {
				rs2 = append(rs2, i)
			}
		}
		fmt.Println(len(rs1)*len(rs1)*len(rs1) + len(rs2)*len(rs2)*len(rs2))
	} else {
		var rs []int
		for i := 1; i <= n; i++ {
			if i%k == 0 {
				rs = append(rs, i)
			}
		}
		fmt.Println(len(rs) * len(rs) * len(rs))
	}
}
