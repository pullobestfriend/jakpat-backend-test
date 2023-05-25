package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal2(t *testing.T) {
	var tests = []struct {
		TestName       string
		Input          string
		ExpectedOutput []string
	}{
		{"found from user on multiple room", "Ari", []string{"backend", "frontend"}},
		{"found from message on single room", "lorem", []string{"frontend"}},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			result := soal2(test.Input)
			assert.Equal(t, test.ExpectedOutput, result)
		})
	}
}
