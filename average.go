package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// averageNumber function that returns the average number of all the numbers in the given string
// difficulty: Medium
// estimated time: 30 minutes
// actual time: 25 minutes
func averageNumber(input string) float64 {
	var average float64
	if testValidity(input) { // verify is input string comply on the format or not
		re := regexp.MustCompile(`\d[\d,]*[\.]?[\d{2}]*`) // regular expression used to separate the numbers from the string
		separatedNumbers := re.FindAllString(input, -1)   // it will give the array of separated numbers
		for _, num := range separatedNumbers {
			number, _ := strconv.ParseUint(num, 10, 32) // covert string to uint
			finalIntNum := int(number)                  //Convert uint64 To int
			average += float64(finalIntNum)
		}

		average /= float64(len(separatedNumbers)) // calculating average
		return average
	}
	fmt.Println("invalid input string")
	return 0
}
