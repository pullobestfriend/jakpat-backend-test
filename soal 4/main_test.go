package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal4(t *testing.T) {
	var tests = []struct {
		TestName       string
		Input          int
		ExpectedOutput []string
		ExpectedErr    error
	}{
		{"valid test - 1", 2023, []string{"21 Jul 2023", "13 Oct 2023"}, nil},
		{"valid test - 2", 2024, []string{"05 Jan 2024", "29 Mar 2024", "21 Jun 2024", "13 Sep 2024", "06 Dec 2024"}, nil},
		{"past year", 2000, nil, errors.New("input cannot be before 2023")},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			result, err := soal4(test.Input)
			assert.Equal(t, test.ExpectedOutput, result)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}
