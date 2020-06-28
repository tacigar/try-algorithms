package math

import (
	"fmt"
	"os"
)

func Min(ns ...int) int {
	if len(ns) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "Arguments are necessary")
		return 0
	}
	if len(ns) == 1 {
		return ns[0]
	}
	min := ns[0]
	for i := 1; i < len(ns); i++ {
		if min > ns[i] {
			min = ns[i]
		}
	}
	return min
}

func Max(ns ...int) int {
	if len(ns) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "Arguments are necessary")
		return 0
	}
	if len(ns) == 1 {
		return ns[0]
	}
	max := ns[0]
	for i := 1; i < len(ns); i++ {
		if max < ns[i] {
			max = ns[i]
		}
	}
	return max
}
