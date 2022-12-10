package day4

import (
	"testing"
)


func TestGetSectonAssignmentOverlap(t *testing.T)  {
	// Test cases
	test_cases := []struct {
		strs   []string
		expected int
	}{
		{
			strs:[]string{ "2-4", "6-8", "2-3",
			"4-5", "5-7", "7-9", "2-8",
			"3-7", "6", "4-6", "2-6",
			"4-8"},
			expected: 2, 
	},
	}

	for _, test_case := range test_cases {
		actual := GetSectionAssignmentOverlap(test_case.strs)
		if actual != test_case.expected {
			t.Errorf("Expected %d, got %d", test_case.expected, actual)
		}
	}

}
