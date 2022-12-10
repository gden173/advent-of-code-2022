package day3

import (
	"advent-of-go/day1"
	"fmt"
	"strings"
)

// Solution to advent of code day 3
func GetRucksackPriority(strs []string) int  {

	letter_map := make(map[string]int)
	letters := make([]string,0)

	for ch := 'a'; ch <= 'z'; ch++ {
		letters = append(letters, string(ch))
	} 

	for ch := 'A'; ch <= 'Z'; ch++ {
		letters = append(letters, string(ch))
	}
	for i, s := range letters {
		letter_map[s] = i + 1
	}
	scores := 0
	for _, s  := range strs {

		n := len(s)/2
		compartment_one := s[:n]
		compartment_two := s[n:]

		if len(compartment_one) != len(compartment_two) {
			panic("Compartment one and two are not the same length")
		}

		for _, c := range compartment_one {
			if strings.Contains(compartment_two, string(c)) {
				scores += letter_map[string(c)]
				break
			}
		}
	}
	return scores
}

func RunDayThreePartOne(filename string)  {
	strs := day1.ReadFileToArray(filename)
	fmt.Println("Day3 Part 1: ", GetRucksackPriority(strs))
}


// Day 3 part two


func GetRucksackPriorityElfGroup(strs []string) int {
	letter_map := make(map[string]int)
	letters := make([]string,0)

	for ch := 'a'; ch <= 'z'; ch++ {
		letters = append(letters, string(ch))
	} 

	for ch := 'A'; ch <= 'Z'; ch++ {
		letters = append(letters, string(ch))
	}
	for i, s := range letters {
		letter_map[s] = i + 1
	}
	scores := 0
	for i := 0; i < len(strs); i+=3 {
		compartment_one := strs[i]
		compartment_two := strs[i+1]
		compartment_three := strs[i+2]

		// create a map containing the unique characters in each compartment 
		for _, c := range compartment_one {
			if strings.Contains(compartment_two, string(c)) && strings.Contains(compartment_three, string(c)) {
				scores += letter_map[string(c)]
				break
			}	
		}
	}
	return scores
}


func RunDayThreePartTwo(filepath string) {
	strs := day1.ReadFileToArray(filepath)
	fmt.Println("Day 3 Part 2: ", GetRucksackPriorityElfGroup(strs))
}
