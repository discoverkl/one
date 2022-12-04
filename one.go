package one

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Pick first non-empty value.
func Pick[T comparable](args ...T) T {
	var zero T
	for _, v := range args {
		if v != zero {
			return v
		}
	}
	return zero
}