package arithmetic

import "testing"

func TestAddSignedOperands(t *testing.T) {
	tests := []struct {
		name string
		left int
		right int
		want int
	}{
		{name: "both negative", left: -2, right: -3, want: -5},
		{name: "mixed signs", left: -2, right: 3, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.left, tt.right); got != tt.want {
				t.Fatalf("Add(%d, %d) = %d; want %d", tt.left, tt.right, got, tt.want)
			}
		})
	}
}
