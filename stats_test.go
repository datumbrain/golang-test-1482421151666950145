package main

import "testing"

// arg1 means argument 1 and the expected stands for the 'result we expect'
type addStatsTest struct {
	arg1                  string
	expectedShortestWord  string
	expectedLongestWord   string
	expectedAverageLength float64
	expectedList          []string
}

var addStatsTests = []addStatsTest{
	addStatsTest{"23-abc-48-cabab-55-haha", "abc", "cabab", 4, []string{"haha"}},
	addStatsTest{"1-hello-2-world", "hello", "hello", 5, []string{"hello", "world"}},
	addStatsTest{"1-47959q", "", "", 0, nil},
}

func TestStoryStatsFunctionality(t *testing.T) {
	for _, test := range addStatsTests {
		shortestWord, longestWord, averageLength, list := storyStats(test.arg1)
		if shortestWord != test.expectedShortestWord {
			t.Errorf("Output %v not equal to expected %v", shortestWord, test.expectedShortestWord)
		}
		if longestWord != test.expectedLongestWord {
			t.Errorf("Output %v not equal to expected %v", longestWord, test.expectedLongestWord)
		}
		if averageLength != test.expectedAverageLength {
			t.Errorf("Output %v not equal to expected %v", averageLength, test.expectedAverageLength)
		}
		for i, word := range list {
			if word != test.expectedList[i] {
				t.Errorf("Output %v not equal to expected %v", word, test.expectedList[i])
			}
		}
	}
}
