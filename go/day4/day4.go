package day4

import (
	"advent-of-go/day1"
	"fmt"
	"strconv"
)

// Solves advent of code day 4 part a
func GetSectionAssignmentOverlap(strs []string) int {
	numOverlaps := 0
	for i := 0; i < len(strs); i+=2 {
		// Parse out the .. frome each elfs assignments and convert them to maps of integers  
		elf_one := make(map[int]bool)
		elf_two := make(map[int]bool)

		for _, c := range strs[i] {
			ic, err := strconv.Atoi(string(c))
			if err != nil { 
				continue

			} else {
				elf_one[ic] = true
			}
		}

		for _, c := range strs[i+1] {
			ic, err := strconv.Atoi(string(c))
			if err != nil { 
				continue

			} else {
				elf_two[ic] = true
			}
		}

		// Compare the two maps and count the number of overlaps
		overlap := 0
		for k := range elf_one {
			if elf_two[k] {
				overlap++
			}
		}
		// If the number of overlaps is equaal to the number of keys then its an overlap 
		if overlap >=  len(elf_one) || overlap >= len(elf_two) {
			numOverlaps ++ 
		}

	}
	return numOverlaps
}


func RunDayFourPartOne(filename string) {
	strs := day1.ReadFileToArray(filename)
	fmt.Println("Day 4 part a: ", GetSectionAssignmentOverlap(strs))
}
