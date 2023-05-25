package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func soal4(input int) ([]string, error) {
	// the LCM of 12, 14, and 6 is 84

	// set the start date as a variable using time package
	startDate := time.Date(2023, time.April, 28, 0, 0, 0, 0, time.Local)

	// set the output as list of string
	var output []string

	// sanitize input, cannot be before 2023
	if input < 2023 {
		return nil, errors.New("input cannot be before 2023")
	}

	for startDate.Year() <= input {
		startDate = startDate.AddDate(0, 0, 84)
		if startDate.Year() == input {
			output = append(output, startDate.Format("02 Jan 2006"))
		}
	}
	return output, nil
}

func main() {
	result, _ := soal4(2023)
	fmt.Println(strings.Join(result, ", "))
	result2, _ := soal4(2024)
	fmt.Println(strings.Join(result2, ", "))
}
