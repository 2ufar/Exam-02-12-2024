package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// GetText := bufio.NewScanner(file)

	var numbers []int
	var words []string
	var sum int
	
	// for GetText.Scan() {
	// 	numbers = append(numbers, GetText.Text())
	// 	words = append(words, GetText.Text())
	// 	sum += numbers
	// }
	bubbleSort(numbers)

	numbersFile, err := os.Create("numbers.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer numbersFile.Close()

	for _, num := range numbers {
		fmt.Println(numbersFile, num)
	}

	fmt.Println(numbersFile, "Sum of all numbers:", sum)

	err = numbersFile.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
		return
	}

	wordsFile, err := os.Create("words.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer wordsFile.Close()

	for _, word := range words {
		fmt.Println(wordsFile, word)
	}
}

func vowelChecker(character rune) {
	switch character {
	case 'a', 'A', 'O', 'o':
		fmt.Printf("The provided character %c is a vowel\n", character)
	default:
		fmt.Printf("The provided character %c is a consonant\n", character)
	}
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
