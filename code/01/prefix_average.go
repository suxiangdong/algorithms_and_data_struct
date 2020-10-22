package code

// O(n^2)
func PrefixAverage1(S []int) []float32 {
	n := len(S)
	A := make([]float32, n)

	for i := 0; i < n; i++ {
		total := 0
		for j := 0; j <= i; j++ {
			total += S[j]
		}
		A[i] = float32(total) / float32(i + 1)
	}

	return A
}

// O(n)
func PrefixAverage2(S []int) []float32 {
	total := 0
	A := make([]float32, len(S))

	for i, v := range S {
		total += v
		A[i] = float32(total) / float32(i+1)
	}

	return A
}