package main

import (
	"advent-of-go/day1"
	"advent-of-go/day2"
	"advent-of-go/day3"
	"advent-of-go/day4"
)

func main() {
	// Day 1 part one
	day1.RunDayOnePartOne("day1/day1a.txt")

	// Day 2 part two
	day1.RunDayOnePartTwo("day1/day1a.txt")

	// Day 2 part one
	day2.RunDayTwoPartOne("day2/day2a.txt")

	// Day 2 part two
	day2.RunDayTwoPartTwo("day2/day2a.txt")

	// Day 3 part one
	day3.RunDayThreePartOne("day3/day3a.txt")

	// Day 3 part two
	day3.RunDayThreePartTwo("day3/day3a.txt")

	// Day 4 part one
	day4.RunDayFourPartOne("day4/day4a.txt")

	// Day 4 part two 
	day4.RunDayFourPartTwo("day4/day4a.txt")

}
