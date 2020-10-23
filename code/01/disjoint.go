package code

// O(n^3)
func Disjoint1(A, B, C []int) bool {
	for _, a := range A {
		for _, b := range B {
			for _, c := range C {
				if a == b && b == c {
					return false
				}
			}
		}
	}

	return true
}

// O(n^2)
func Disjoint2(A, B, C []int) bool {
	for _, a := range A {
		for _, b := range B {
			if a == b {
				for _, c := range C {
					if b == c {
						return false
					}
				}
			}
		}
	}

	return true
}
