package arithmetic

import "testing"

func TestWithinDistanceZero(t *testing.T) {
	if !WithinDistance(0, 0, 0) {
		t.Fatal("WithinDistance(0, 0, 0) = false; want true")
	}
}

func TestWithinDistance(t *testing.T) {
	if !WithinDistance(2, 7, 5) {
		t.Fatal("WithinDistance(2, 7, 5) = false; want true")
	}
}
