package day1

import "testing"

func Test_DayOnePartOne(t *testing.T) {

	test_input := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	expected := 24000

	ret := DayOnePartOne(test_input)

	if ret != expected {
		t.Fail()
	}
}

func Test_DayOnePartTwo(t *testing.T) {

	test_input := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	expected := 45000

	ret := DayOnePartTwo(test_input)

	if ret != expected {
		t.Fail()
	}
}
