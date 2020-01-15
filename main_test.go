package main

import (
	"testing"
)

func TestPlateCalc(t *testing.T) {
	tables := []struct {
		input  float64
		output []float64
	}{
		{0, []float64{}},
		{45, []float64{}},
		{50, []float64{2.5}},
		{95, []float64{25}},
		{110, []float64{25, 5, 2.5}},
	}

	for _, table := range tables {
		actual := plateCalc(table.input)

		if !equal(actual, table.output) {
			t.Errorf("did not get expected result: %+v != %+v", table.output, actual)
		}
	}
}

func equal(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
