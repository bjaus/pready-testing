package arithmetic

import "testing"

func TestAddSigned(t *testing.T) {
	if got := Add(-2, 3); got != 1 {
		t.Fatalf("Add(-2, 3) = %d; want 1", got)
	}
}
