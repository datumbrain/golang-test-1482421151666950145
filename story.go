package main

import (
	"regexp"
	"strings"
)

// wholeStory function that takes the string, and returns a text that is composed of all the text words separated by spaces
// difficulty: Medium
// estimated time: 30 minutes
// actual time: 28 minutes
func wholeStory(input string) string {
	var wholeStory string
	if testValidity(input) { // verify is input string comply on the format or not
		re := regexp.MustCompile(`([a-z]+)`)           // regular expression used to separate the text words from the string
		separatedWords := re.FindAllString(input, -1)  // it will give the array of separated text words
		wholeStory = strings.Join(separatedWords, " ") // it will join the words to make a story string
		return wholeStory
	}
	return ""
}
