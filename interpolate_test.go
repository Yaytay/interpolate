package main_test

import (
	"testing"
	"strings"
	"fmt"
)

func TestInterpolate(t *testing.T) {
	tests := []struct {
		input string
		output string
	}{
		{"", ""},
		{"nan", "NaN"},
		{"1, 2", "1,2"},
		{"1, 2\n3,4", "1,2\n3,4"},
		{"1, nan, 3", "1,2,3"},
		{"1, nan, 3\n4, nan, 6", "1,2,3\n4,5,6"},
		{"1, 2, 3\n4, nan, 6", "1,2,3\n4,4,6"},
		{"59.865848,nan\n60.111501,70.807258",
			"59.865848,65.336553\n60.111501,70.807258"},
		{"37.454012,95.071431,73.199394,59.865848,nan\n15.599452,5.808361,86.617615,60.111501,70.807258\n2.058449,96.990985,nan,21.233911,18.182497\nnan,30.424224,52.475643,43.194502,29.122914\n61.185289,13.949386,29.214465,nan,45.606998",
			"37.454012,95.071431,73.199394,59.865848,65.336553\n15.599452,5.808361,86.617615,60.111501,70.807258\n2.058449,96.990985,64.3295385,21.233911,18.182497\n31.222654,30.424224,52.475643,43.194502,29.122914\n61.185289,13.949386,29.214465,39.338655,45.606998"},
	}
	for _, test := range tests {
		rdr := strings.NewReader(test.input)
		var wrtr strings.Builder
		Interpolate(rdr, &wrtr)
		result := wrtr.String()
		fmt.Printf("%v", result)
		if result != test.output {
			t.Errorf("%v should have returned '%v', not '%v'", test.input, test.output, result)
		}
	}
}