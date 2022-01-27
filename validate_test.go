package main

import "testing"

// arg1 means argument 1 and the expected stands for the 'result we expect'
type addValidityTest struct {
	arg1     string
	expected bool
}

var addValidityTests = []addValidityTest{
	addValidityTest{"23-ab-48-caba-56-haha", true},
	addValidityTest{"1-hello-2-world", true},
	addValidityTest{"6-e-6-i-7-yv-38-b", true},
	addValidityTest{"1-47959q", false},
}

func TestValidateFunctionality(t *testing.T) {
	for _, test := range addValidityTests {
		if output := testValidity(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
