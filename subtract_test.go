package arithmetic

import "testing"

func TestSubtract(t *testing.T) {
	if got := Subtract(7, 2); got != 5 {
		t.Fatalf("Subtract(7, 2) = %d; want 5", got)
	}
}
