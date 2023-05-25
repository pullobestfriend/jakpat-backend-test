package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal3(t *testing.T) {
	var tests = []struct {
		TestName       string
		Input          int
		ExpectedOutput string
		ExpectedErr    error
	}{
		{"valid test - 1", 3, "From first page flip 1x", nil},
		{"valid test - 2", 4, "From first page flip 2x", nil},
		{"valid test - 3", 230, "From last page flip 0x", nil},
		{"out of bounds - 1", 256, "", errors.New("input is out of range")},
		{"out of bounds - 2", -1, "", errors.New("input is out of range")},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			result, err := soal3(test.Input)
			assert.Equal(t, test.ExpectedOutput, result)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}
