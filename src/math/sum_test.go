package math

import "testing"

func TestSum(t *testing.T) {
	result := Sum(1, 2)
	if result != 3 {
		t.Error("Expect 3, but got ", result)
	}
}
