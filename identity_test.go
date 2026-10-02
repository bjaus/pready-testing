package arithmetic

import "testing"

func TestAddZeroIdentity(t *testing.T) {
	if got := Add(7, 0); got != 7 {
		t.Fatalf("Add(7, 0) = %d; want 7", got)
	}
}
