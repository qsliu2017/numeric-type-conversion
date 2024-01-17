package numeric_type_conversion_test

import (
	"math"
	"testing"
)

func float64ToUint64(f float64) uint64 {
	return uint64(f)
}

func TestConversions(t *testing.T) {
	for _, tc := range []struct {
		input    interface{}
		name     string
		expected interface{}
	}{
		{input: float64ToUint64(-2.0), name: "uint64(-2.0)", expected: uint64(math.MaxUint64) - 1},
	} {
		if tc.input != tc.expected {
			t.Errorf("%s expected %v, got %v", tc.name, tc.expected, tc.input)
		}
	}
}
