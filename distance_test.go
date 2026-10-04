package arithmetic

import "testing"

func TestDistanceZero(t *testing.T) {
	if got := Distance(0, 0); got != 0 {
		t.Fatalf("Distance(0, 0) = %d; want 0", got)
	}
}

func TestDistance(t *testing.T) {
	if got := Distance(2, 7); got != 5 {
		t.Fatalf("Distance(2, 7) = %d; want 5", got)
	}
}
