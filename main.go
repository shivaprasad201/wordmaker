package main

import (
	"fmt"
	"strings"

	"./helpers"
)

func main() {
	var letterCount int
	var wordLength int

	fmt.Println("Enter the number of letters")
	fmt.Scan(&letterCount)

	fmt.Println("Enter the word length")
	fmt.Scan(&wordLength)

	var lettersArray []string

	//Reading to the array
	for i := 0; i < letterCount; i++ {
		letterPosition := i + 1
		var readLetter string

		fmt.Println("Enter ", letterPosition, "th letter")
		fmt.Scan(&readLetter)

		lettersArray = append(lettersArray, readLetter)
	}

	allLetterString := strings.Join(lettersArray[:], "")

	//Calling word generator
	combinationArray := helpers.WordGen(allLetterString, wordLength)

	for _, word := range combinationArray {
		definition, wordValidity := helpers.IsWord(word)
		if wordValidity {
			fmt.Println("..........")
			fmt.Println(word)
			fmt.Println("Definition is \n", definition)
			fmt.Println("..........")
		}

	}
}
