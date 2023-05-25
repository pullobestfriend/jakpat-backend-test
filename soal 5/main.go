package main

import (
	"fmt"
	"math"
	"strings"
)

func soal5(input string) string {
	// trim whitespace
	trimmed := strings.ReplaceAll(input, " ", "")

	// find square root from the lenght of input
	root := math.Sqrt(float64(len(trimmed)))

	// find the row and column of root
	row := int(math.Floor(root))
	col := int(math.Ceil(root))

	// create array 2D of the characters
	arr := make([][]string, row)
	for i := range arr {
		arr[i] = make([]string, col)
	}

	// insert data into array
	n := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			arr[i][j] = fmt.Sprintf("%c", trimmed[n])
			n++
			// if the input cannot fully fill the array, need to break
			if n == len(trimmed) {
				break
			}
		}
	}

	var output string
	// transpose array
	for j := 0; j < col; j++ {
		for i := 0; i < row; i++ {
			output += arr[i][j]
		}
		// use if to eliminate extra whitespace
		if j < col-1 {
			output += " "
		}
	}
	return output
}

func main() {
	fmt.Println(soal5("aku wibu dan aku bangga"))
	fmt.Println(soal5("have a nice day"))
	fmt.Println(soal5("a"))
	fmt.Println(soal5(""))
}
