package main

import (
	"fmt"
	"math"
	"regexp"
)

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
