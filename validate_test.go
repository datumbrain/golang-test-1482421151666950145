package main

import "testing"

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
	arg1     string
	expected bool
}

var addTests = []addTest{
	addTest{"23-ab-48-caba-56-haha", true},
	addTest{"1-hello-2-world", true},
	addTest{"6-e-6-i-7-yv-38-b", true},
	addTest{"1-47959q", false},
}

func TestValidate(t *testing.T) {
	for _, test := range addTests {
		if output := testValidity(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
