package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func scan() string {
	scanner.Scan()
	return scanner.Text()
}

const N = 10

func main() {
	scanner.Split(bufio.ScanWords)
	f := [][]byte{}
	for i := 0; i < N; i++ {
		f = append(f, []byte(scan()))
	}
	var a, b, c, d = -1, -1, -1, -1
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if f[i][j] == '#' {
				a = i
				c = j
				for k := j + 1; k < N; k++ {
					if f[i][k] == '.' {
						d = k - 1
						break
					}
				}
				if d == -1 {
					d = N - 1
				}
				for k := i + 1; k < N; k++ {
					if f[k][j] == '.' {
						b = k - 1
						break
					}
				}
				if b == -1 {
					b = N - 1
				}
				goto loopEnd
			}
		}
	}
loopEnd:
	fmt.Printf("%d %d\n%d %d\n", a+1, b+1, c+1, d+1)
}
