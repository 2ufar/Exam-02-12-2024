package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Employee struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Salary     int    `json:"salary"`
	Experience int    `json:"experience"`
	Age        int    `json:"age"`
}


func main() {
	jsonData, err := os.ReadFile("employees.json")
	if err != nil {
		log.Fatal(err)
	}
	
	var employees []Employee
	err = json.Unmarshal(jsonData, &employees)
	if err != nil {
		log.Fatal(err)
	}
	
	bubbleSort(employees)
	
	file, err := os.Create("orderByAge.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	for _, employee := range employees {
		err = json.NewEncoder(file).Encode(employee)
	}
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("File created successfully:", "orderByAge.txt")
	
	bubbleSort(employees)
	
	file2, err2 := os.Create("topSalaryEmployees.txt")
	if err2 != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	
	for i := 0; i < 3; i++ {
		err = json.NewEncoder(file2).Encode(employees[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("File created successfully:", "topSalaryEmployees.txt")
	
}

func bubbleSort(employees []Employee) {
	n := len(employees)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if employees[j].Age > employees[j+1].Age {
				employees[j], employees[j+1] = employees[j+1], employees[j]
			}
		}
	}
}