package main

import (
	"errors"
	"fmt"
	"strings"
)

func soal1(input string) (string, error) {
	// remove any -
	trimmed := strings.ReplaceAll(input, "-", "")

	// check first digit, if not 0 add 0
	if !strings.HasPrefix(trimmed, "0") {
		trimmed = "0" + trimmed
	}

	// check total digit, if more than 12 return error
	if len(trimmed) > 12 {
		return "", errors.New("digit is more than 12")
	}

	// add - after every 4 digit
	// in case the number have 10 or 11 digit, it will be formatted as 0xxx-xxxx-xx or 0xxx-xxxx-xxx
	output := trimmed[:4] + "-" + trimmed[4:8] + "-" + trimmed[8:]

	return output, nil
}

func main() {
	fmt.Println(soal1("81234567890"))
	fmt.Println(soal1("081-234-567-890"))
}
