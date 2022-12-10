package day4

import "testing"


func TestGetSectonAssignmentOverlap(t *testing.T)  {
	// Test cases
	test_cases := []struct {
		strs   []string
		expected int
	}{
		{
			strs:[]string{ ".234.....", ".....678.", ".23......", "...45....", "....567..", "......789", ".2345678.", "..34567..", ".....6...", "...456...", ".23456...", "...45678."},
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
