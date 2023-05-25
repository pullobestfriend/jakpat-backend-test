package main

import (
	"errors"
	"fmt"
)

func soal3(input int) (string, error) {
	// sanitize input, input must be 1 <= x <= 231
	if input < 1 || input > 231 {
		return "", errors.New("input is out of range")
	}

	// set the middle ground, (231-1)/2 = 115
	// if <= 115, go from first page, else go from last page
	if input <= 115 {
		output := input / 2
		return fmt.Sprint("From first page flip ", output, "x"), nil
	} else {
		output := (231 - input) / 2
		return fmt.Sprint("From last page flip ", output, "x"), nil
	}
}

func main() {
	fmt.Println(soal3(3))
	fmt.Println(soal3(4))
	fmt.Println(soal3(230))
}
