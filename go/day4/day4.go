package day4

import (
	"advent-of-go/day1"
	"fmt"
	"strconv"
	"strings"
)

func stringToRange(s string) map[int]bool {
	// Parse the into into two
	ret := make(map[int]bool)

	if !strings.Contains(s, "-") {
		val, err := strconv.Atoi(strings.Trim(s, " "))
		if err != nil {
			panic("Cannot convert to integer")
		}
		ret[val] = true
		return ret
	}

	st := strings.Split(s, "-")

	if len(st) != 2 {
		fmt.Println(st)
		panic("Too many items in string: should only be 2")
	}

	s1, _ := strconv.Atoi(strings.Trim(st[0], " "))

	s2, _ := strconv.Atoi(strings.Trim(st[1], " "))

	for i := s1; i < s2+1; i++ {
		ret[i] = true
	}

	return ret
}

// Solves advent of code day 4 part a
func GetSectionAssignmentOverlap(strs []string) int {
	numOverlaps := 0
	for i := 0; i < len(strs); i++ {

		s := strings.Split(strs[i], ",")

		elf_one := stringToRange(s[0])
		elf_two := stringToRange(s[1])

		overlap := 0
		for k := range elf_one {
			if elf_two[k] {
				overlap++
			}
		}

		// If the number of overlaps is equaal to the number of keys then its an overlap
		if overlap >= len(elf_one) || overlap >= len(elf_two) {
			numOverlaps++
		}

	}
	return numOverlaps
}

func RunDayFourPartOne(filename string) {
	strs := day1.ReadFileToArray(filename)
	fmt.Println("Day 4 part a: ", GetSectionAssignmentOverlap(strs))
}

// Day 4 part B
func GetTotalAssignmentOverlap(strs []string) int {
	numOverlaps := 0
	for i := 0; i < len(strs); i++ {

		s := strings.Split(strs[i], ",")

		elf_one := stringToRange(s[0])
		elf_two := stringToRange(s[1])

		overlap := 0
		for k := range elf_one {
			if elf_two[k] {
				overlap++
			}
		}

		if overlap > 0 {
			numOverlaps++
		}

	}
	return numOverlaps
}

func RunDayFourPartTwo(filename string) {
	strs := day1.ReadFileToArray(filename)
	fmt.Println("Day 4 part B: ", GetTotalAssignmentOverlap(strs))
}
