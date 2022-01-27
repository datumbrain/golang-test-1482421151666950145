package main

import "testing"

// arg1 means argument 1 and the expected stands for the 'result we expect'
type addStoryTest struct {
	arg1     string
	expected string
}

var addStoryTests = []addStoryTest{
	addStoryTest{"23-ab-48-caba-55-haha", "ab caba haha"},
	addStoryTest{"1-hello-2-world", "hello world"},
	addStoryTest{"12-i-22-am-23-golang-45-developer", "i am golang developer"},
	addStoryTest{"1-47959q", ""},
}

func TestWholeStoryFunctionality(t *testing.T) {
	for _, test := range addStoryTests {
		if output := wholeStory(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
