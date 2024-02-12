package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)


func main() {
	var filePath string
	fmt.Print("Enter file path: ")
	fmt.Scan(&filePath)
	
	for checkFileExists(filePath + ".txt") {
		var numSuggestions int
		fmt.Print("File already exists. Enter the number of suggestions you would like: ")
		fmt.Scan(&numSuggestions)
		
		suggestNames(filePath, numSuggestions)
		
		fmt.Print("Enter a new file path: ")
		fmt.Scan(&filePath)
	}
	
	file, err := os.Create(filePath + ".txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	
	fmt.Println("File created successfully:", filePath+".txt")
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func generateRandomFile(orgName string) string {
	return orgName + strconv.Itoa(rand.Intn(100)) + ".txt"
}

func suggestNames(OrgName string, numSuggestions int) {
	fmt.Println("File already exists. You can create files with names like these:")
	for i := 0; i < numSuggestions; i++ {
		filePath := generateRandomFile(OrgName)
		if !checkFileExists(filePath) {
			fmt.Printf("%v ", filePath)
		}
	}
	fmt.Println()
}