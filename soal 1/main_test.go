package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal1(t *testing.T) {
	var tests = []struct {
		TestName       string
		Input          string
		ExpectedOutput string
		ExpectedErr    error
	}{
		{"no prefix zero", "81234567890", "0812-3456-7890", nil},
		{"misaligned dash", "081-234-567-890", "0812-3456-7890", nil},
		{"with non digit", "081-abc-567-890", "", errors.New("some error")},
		{"with not sufficient numbers", "8123", "", errors.New("some error")},
		{"with too much numbers", "81234567890123456789", "", errors.New("digit is more than 12")},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			result, err := soal1(test.Input)
			assert.Equal(t, test.ExpectedOutput, result)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}
