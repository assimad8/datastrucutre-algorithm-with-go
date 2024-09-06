package divideandconquer

// fibonacci method given k integer
func Fibonacci(k int) int {
	if k<=1 {
		return 1
	}
	return Fibonacci(k-1)+Fibonacci(k-2)
}

