package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// testValidity function returns true or false based on the input string, is it comply on the defined format or not
// difficulty: Easy
// estimated time: 20 minutes
// actual time: 15 minutes
func testValidity(input string) bool {
	format, _ := regexp.Compile("^(^([0-9]+)-([a-z]+)((\\-?)([0-9]+)-([a-z]+))*)$") // defined regular expression to validate the input string
	return format.MatchString(input)                                                // return true if input comply on format and return false if not
}

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

// wholeStory function that takes the string, and returns a text that is composed of all the text words separated by spaces
// difficulty: Medium
// estimated time: 30 minutes
// actual time:
func wholeStory(input string) string {

}

func main() {
	fmt.Println(testValidity("23-ab-48-caba-56-haha")) // calling testValidity function and printing the output

	fmt.Println(averageNumber("1-hello-2-world-10-aa-02-cc")) // calling averageNUmber function and printing the output that must be equal to (1+2+10+2)/4 = 3.75

}
