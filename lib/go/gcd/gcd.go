package gcd

func GCD(a, b int) int {
	if a%b == 0 {
		return b
	}
	return GCD(b, a%b)
}
