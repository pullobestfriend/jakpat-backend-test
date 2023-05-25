package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal5(t *testing.T) {
	var tests = []struct {
		TestName       string
		Input          string
		ExpectedOutput string
	}{
		{"valid test - 1", "aku wibu dan aku bangga", "aban kukg udug waba ina"},
		{"valid test - 2", "have a nice day", "hae and via ecy"},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			result := soal5(test.Input)
			assert.Equal(t, test.ExpectedOutput, result)
		})
	}
}
