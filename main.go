package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
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
// actual time: 28 minutes
func wholeStory(input string) string {
	var wholeStory string
	if testValidity(input) { // verify is input string comply on the format or not
		re := regexp.MustCompile(`([a-z]+)`)           // regular expression used to separate the text words from the string
		separatedWords := re.FindAllString(input, -1)  // it will give the array of separated text words
		wholeStory = strings.Join(separatedWords, " ") // it will join the words to make a story string
		return wholeStory
	}
	fmt.Println("invalid input string")
	return ""
}

// storyStats function that returns the shortest word, the longest word, average word length and the list (or empty list) of all words
// from the story that have same length as the average length rounded up and down.
// difficulty: Medium
// estimated time: 50 minutes
// actual time: 1 hour and 4 minutes
func storyStats(input string) (string, string, float64, []string) {
	var list []string
	var shortestWord, longestWord string
	var averageLen float64
	if testValidity(input) { // verify is input string comply on the format or not
		re := regexp.MustCompile(`([a-z]+)`)          // regular expression used to separate the text words from the string
		separatedWords := re.FindAllString(input, -1) // it will give the array of separated text words
		shortestWord = separatedWords[0]
		longestWord = separatedWords[0]
		for _, word := range separatedWords { // traverse int the separatedWords array
			if len(word) > len(longestWord) { // if the current word is larger than the previous one, save it in the largestWord variable
				longestWord = word
			}
			if len(word) < len(shortestWord) { // if the current word is shorter than the previous one, save it in the shortestWord variable
				shortestWord = word
			}

			averageLen += float64(len(word))
		}

		averageLen /= float64(len(separatedWords))      // calculating average length of the words
		roundedUpAverageLen := math.Ceil(averageLen)    // rounded up
		roundedDownAverageLen := math.Floor(averageLen) // rounded down
		for _, word := range separatedWords {           // find and store the words whose length equals to rounded average length
			if float64(len(word)) == roundedUpAverageLen || float64(len(word)) == roundedDownAverageLen {
				list = append(list, word)
			}
		}

		return shortestWord, longestWord, averageLen, list
	}

	fmt.Println("invalid input string")
	return "", "", 0, nil
}

// generate function, takes boolean flag and generates random correct strings if the parameter is true and random
// incorrect strings if the flag is false.
// difficulty: medium
// estimated time: 1 hour and 30 minutes
// actual time: 1 hour and 25 minutes
func generate(flag bool) string {

}

func main() {
	// calling testValidity function and printing the output
	isValid := testValidity("23-ab-48-caba-56-haha")
	fmt.Println(isValid)
	fmt.Println() // printing empty line

	// calling averageNUmber function and printing the output that must be equal to (1+2+10+2)/4 = 3.75
	average := averageNumber("1-hello-2-world-10-aa-02-cc")
	fmt.Printf("average of the numbers = %v\n", average)
	fmt.Println() // printing empty line

	// calling the wholeStory function and printing the output
	story := wholeStory("1-hello-2-world")
	fmt.Printf("story: %v\n", story)
	fmt.Println() // printing empty line

	// calling storyStats function and printing the output
	shortestWord, longestWord, averageLength, list := storyStats("11-hello-22-hi-33-world-44-golang-55-task")
	fmt.Printf("shortest word: %s\n", shortestWord)
	fmt.Printf("longest word: %s\n", longestWord)
	fmt.Printf("average word length: %v\n", averageLength)
	fmt.Printf("list of the words having length equals to the rounded average length: %v\n", list)
	fmt.Println() // printing empty line
}
