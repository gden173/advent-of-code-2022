package day3

import (
	"testing"
)

func TestGetRucksackPriority(t *testing.T) {
	// Test cases
	test_cases := []struct {
		strs     []string
		expected int
	}{
		{
			[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
			157,
		},
	}

	for _, test_case := range test_cases {
		actual := GetRucksackPriority(test_case.strs)
		if actual != test_case.expected {
			t.Errorf("Expected %d, got %d", test_case.expected, actual)
		}
	}
}

func TestGetRucksackPriorityElfGroups(t *testing.T) {
	// Test cases
	test_cases := []struct {
		strs     []string
		expected int
	}{
		{
			[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
			70,
		},
	}

	for _, test_case := range test_cases {
		actual := GetRucksackPriorityElfGroup(test_case.strs)
		if actual != test_case.expected {
			t.Errorf("Expected %d, got %d", test_case.expected, actual)
		}
	}
}
