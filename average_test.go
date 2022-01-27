package main

import "testing"

// arg1 means argument 1 and the expected stands for the 'result we expect'
type addAverageTest struct {
	arg1     string
	expected float64
}

var addAverageTests = []addAverageTest{
	addAverageTest{"23-ab-48-caba-55-haha", 42},
	addAverageTest{"1-hello-2-world", 1.5},
	addAverageTest{"6-e-6-i-7-yv-38-b", 14.25},
	addAverageTest{"1-47959q", 0},
}

func TestAverageNumberFunctionality(t *testing.T) {
	for _, test := range addAverageTests {
		if output := averageNumber(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
