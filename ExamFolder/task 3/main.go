package main

import (
	"fmt"
)

type Person struct {
	Name string
	Sex  string
	Job  string
	Age  int
}

func main() {
	var totalPeople int
	fmt.Print("Enter the total number of people: ")
	_, err := fmt.Scan(&totalPeople); 
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	peopleList := make([]Person, 0)

	for i := 0; i < totalPeople; i++ {
		fmt.Printf("\nInformation for person %d:\n", i+1)
		var person Person

		fmt.Print("Name: ")
		_, NameError := fmt.Scan(&person.Name); 
		if NameError != nil {
			fmt.Println("Error:", NameError)
			return
		}

		fmt.Print("Sex (male/female): ")
		_, GenderError := fmt.Scan(&person.Sex)
		if GenderError != nil {
			fmt.Println("Error:", NameError)
			return
		}

		fmt.Print("Job: ")
		_, JobError := fmt.Scan(&person.Job)
		if JobError != nil {
			fmt.Println("Error:", NameError)
			return
		}

		fmt.Print("Age: ")
		_, AgeError := fmt.Scan(&person.Age) 
		if AgeError != nil {
			fmt.Println("Error:", NameError)
			return
		}

		peopleList = append(peopleList, person)
	}

	var category string
	fmt.Print("\nChoose a category (males/females/all): ")
	_, CategoryError := fmt.Scan(&category)
	if CategoryError != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nResults:")
	for _, person := range filterPeople(peopleList, category) {
		fmt.Printf(" Name: %s\n Sex: %s\n Job: %s\n Age: %d\n\n", person.Name, person.Sex, person.Job, person.Age)
	}
}

func filterPeople(peopleList []Person, category string) []Person {
	filteredPeople := make([]Person, 0)

	for _, person := range peopleList {
		switch category {
		case "males":
			if person.Sex == "male" {
				filteredPeople = append(filteredPeople, person)
			}
		case "females":
			if person.Sex == "female" {
				filteredPeople = append(filteredPeople, person)
			}
		case "all":
			filteredPeople = append(filteredPeople, person)
		default:
			fmt.Println("Invalid category!")
			return nil
		}
	}

	return filteredPeople
}