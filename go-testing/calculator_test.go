package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3
	if result != expected {
		t.Errorf("Add(1, 2) failed. Expected 3, got %d", result)
	}
}

type divideTestCase struct {
	a, b           int
	expectedResult int
	expectError    bool
}

func TestDivide(t *testing.T) {
	testCases := []divideTestCase{
		{a: 6, b: 3, expectedResult: 2, expectError: false},
		{a: 6, b: 0, expectedResult: 0, expectError: true},
	}

	for _, tc := range testCases {
		result, err := Divide(tc.a, tc.b)
		if tc.expectError {
			assert.Error(t, err, "an unexpected error occurred: %v", err)
		} else {
			assert.NoError(t, err, "an unexpected error occurred: %v", err)
			assert.Equal(t, tc.expectedResult, result, "they should be equal")
		}
	}
}
