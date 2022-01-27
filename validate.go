package main

import "regexp"

// testValidity function returns true or false based on the input string, is it comply on the defined format or not
// difficulty: Easy
// estimated time: 20 minutes
// actual time: 15 minutes
func testValidity(input string) bool {
	format, _ := regexp.Compile("^(^([0-9]+)-([a-z]+)((\\-?)([0-9]+)-([a-z]+))*)$") // defined regular expression to validate the input string
	return format.MatchString(input)                                                // return true if input comply on format and return false if not
}
