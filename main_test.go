package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},       // Test Case 1
		{-1, 1, 0},      // Test Case 2
		{0, 0, 0},       // Test Case 3
		{10, -5, 5},     // Test Case 4
		{100, 200, 300}, // Test Case 5
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%d+%d=%d", tc.a, tc.b, tc.expected), func(t *testing.T) {
			t.Parallel()
			actual := Add(tc.a, tc.b)
			if actual != tc.expected {
				t.Errorf("expected: %d, actual: %d", tc.expected, actual)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b     int
		expected int
		err      error
	}{
		{10, 2, 5, nil}, // Test Case 1
		{10, 0, 0, fmt.Errorf("cannot divide by zero")}, // Test Case 2
		{0, 0, 0, fmt.Errorf("cannot divide by zero")},  // Test Case 3
		{100, 200, 0, nil}, // Test Case 4
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%d/%d=%d", tc.a, tc.b, tc.expected), func(t *testing.T) {
			t.Parallel()
			actual, err := Divide(tc.a, tc.b)
			if actual != tc.expected {
				t.Errorf("expected: %d, actual: %d", tc.expected, actual)
			}
			if err != nil {
				if tc.err == nil {
					t.Errorf("unexpected error: %v", err)
				} else if err.Error() != tc.err.Error() {
					t.Errorf("expected error: %v, actual: %v", tc.err, err)
				}
			}
		})
	}
}
