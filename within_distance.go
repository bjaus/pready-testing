package arithmetic

func WithinDistance(a, b, limit int) bool {
	return Distance(a, b) <= limit
}
