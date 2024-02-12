package main

import (
	"fmt"
	"os"
)

type Student struct {
	Name   string
	Grades []int
	Course string
}

func (s *Student) calculateAverage() float64 {
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	averageCalculator :=  float64(sum) / float64(len(s.Grades))
	return averageCalculator
}

func main() {
    student := &Student{Name: "John", Grades: []int{70, 80, 23, 90}, Course: "Math"}

    average := student.calculateAverage()

    var result string
    if average < 60 {
        result = "failed"
    } else {
        result = "passed"
    }

    file, err := os.Create("result.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    _,err = file.WriteString(fmt.Sprintf("Student %s %s the course %s with an average grade of %.2f\n", student.Name, result, student.Course, average))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Results have been written to result.txt")
}