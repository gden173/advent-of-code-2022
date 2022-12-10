package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFileToArray(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		panic("error")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Iterate through until a blank line
func DayOnePartOne(strs []string) int {
	currentVal := 0
	maxVal := 0
	for _, s := range strs {
		if s == "" {
			if maxVal < currentVal {
				maxVal = currentVal
			}

			currentVal = 0
			continue
		}

		is, err := strconv.Atoi(s)
		if err != nil {
			panic("Cannot convert int")
		}

		currentVal += is
	}

	return maxVal
}

func RunDayOnePartOne(fileName string) {
	s := ReadFileToArray(fileName)
	res := DayOnePartOne(s)
	fmt.Printf("Day 1: Part A: The Elf with the most calories had %d calories\n", res)

}

// Return the sum of the top three elves using the previous input
func DayOnePartTwo(strs []string) int {
	currentVal := 0
	amounts := make([]int, 0)

	for _, s := range strs {
		if s == "" {
			amounts = append(amounts, currentVal)
			currentVal = 0
			continue
		}

		is, err := strconv.Atoi(s)
		if err != nil {
			panic("Cannot convert int")
		}
		currentVal += is
	}
	amounts = append(amounts, currentVal)

	top3 := [3]int{0, 0, 0}

	for _, is := range amounts {
		if is > top3[0] {
			if is > top3[1] {
				if is > top3[2] {

					top3[0] = top3[1]
					top3[1] = top3[2]
					top3[2] = is
					continue
				}
				top3[0] = top3[1]
				top3[1] = is
				continue

			}
			top3[0] = is
			continue
		}
	}

	return top3[0] + top3[1] + top3[2]
}

func RunDayOnePartTwo(fileName string) {
	s := ReadFileToArray(fileName)
	res := DayOnePartTwo(s)
	fmt.Printf("Day 1: Part B: The sum of top three elves calories is %d\n", res)
}
