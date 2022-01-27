package main

import (
	"fmt"
)

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
	story := wholeStory("12-i-22-am-23-golang-45-developer")
	fmt.Printf("story: %v\n", story)
	fmt.Println() // printing empty line

	// calling storyStats function and printing the output
	shortestWord, longestWord, averageLength, list := storyStats("11-hello-22-hi-33-world-44-golang-55-task")
	fmt.Printf("shortest word: %s\n", shortestWord)
	fmt.Printf("longest word: %s\n", longestWord)
	fmt.Printf("average word length: %v\n", averageLength)
	fmt.Printf("list of the words having length equals to the rounded average length: %v\n", list)
	fmt.Println() // printing empty line

	// calling generate function and printing the output
	str := generate(true)
	fmt.Printf("random generated string: %v\n", str)
}
