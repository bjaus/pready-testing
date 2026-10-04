package arithmetic

func Distance(a, b int) int {
	if a >= b {
		return Subtract(a, b)
	}
	return Subtract(b, a)
}
