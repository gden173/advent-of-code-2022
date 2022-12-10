package day2

import (
	"testing"
)

func TestPaperScissorsRock(t *testing.T) {
	// Test cases
	test_cases := []struct {
		strategyGuide []string
		expected      int
	}{
		{
			[]string{"A Y", "B X", "C Z"},
			15,
		},
	}

	for _, test_case := range test_cases {
		actual := PaperScissorsRock(test_case.strategyGuide)
		if actual != test_case.expected {
			t.Errorf("Expected %d, got %d", test_case.expected, actual)
		}
	}
}

func TestPaperScissorsRockWinLoseDraw(t *testing.T) {
	// Test cases
	test_cases := []struct {
		strategyGuide []string
		expected      int
	}{
		{
			[]string{"A Y", "B X", "C Z"},
			12,
		},
	}

	for _, test_case := range test_cases {
		actual := PaperScissorsRockWinLoseDraw(test_case.strategyGuide)
		if actual != test_case.expected {
			t.Errorf("Expected %d, got %d", test_case.expected, actual)
		}
	}
}
