package main

import (
	"math"
	"strings"
	"testing"
)

func equal(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !math.IsNaN(v) {
			if math.IsNaN(b[i]) {
				return false
			}
			// This suffers from known issues with floating point comparisons
			// As long as it's OK for the unit tests I'll leave it
			// Correct fix requires knowledge of the values being used
			if v != b[i] {
				return false
			}
		}
	}
	return true
}

func TestReadLine(t *testing.T) {
	tests := []struct {
		line   string
		result []float64
	}{
		{"1, 2", []float64{1.0, 2.0}},
		{"1,2,3,4", []float64{1.0, 2.0, 3.0, 4.0}},
		{"1,2,nan,4", []float64{1.0, 2.0, math.NaN(), 4.0}},
		{"1,2,bob,4", []float64{1.0, 2.0, math.NaN(), 4.0}},
		{"1.1,2.2,bob,4.4", []float64{1.1, 2.2, math.NaN(), 4.4}},
		{"1.1,2.2,-inf,4.4", []float64{1.1, 2.2, math.Inf(-1), 4.4}},
	}

	// TODO: Give Interpolate a field that defaults to os.Stderr so tests can capture error messages

	for _, test := range tests {
		rdr := prepCsvReader(strings.NewReader(test.line))
		result, _ := readLine(rdr, 2, 3)
		if !equal(result, test.result) {
			t.Errorf("%v should have returned %v", test.line, test.result)
		}
	}
}
