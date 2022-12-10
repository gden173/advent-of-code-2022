package day2

import (
	"advent-of-go/day1"
	"strings"
)

// Solves advent of code day 2 in golang

func PaperScissorsRock(strategyGuide []string) int {
	// Opponent's strategy
	// A -> Rock
	// B -> Paper
	// C -> Scissors
	// Your stragegy
	// X -> Rock
	// Y -> paper
	// z ->  scissors
	// Create a score map for every combination
	move_map := make(map[string]string)
	move_map["A"] = "rock"
	move_map["X"] = "rock"
	move_map["B"] = "paper"
	move_map["Y"] = "paper"
	move_map["C"] = "scissors"
	move_map["Z"] = "scissors"

	// create a score map for every move
	score_map := make(map[string]int)
	score_map["rock"] = 1
	score_map["paper"] = 2
	score_map["scissors"] = 3

	// Create a wl map fore winning and losing
	wl_map := make(map[string]int)
	wl_map["win"] = 6
	wl_map["draw"] = 3
	wl_map["loss"] = 0

	scores := 0
	for _, strategy := range strategyGuide {
		// Splite the line into your score and the opponent's score
		strs := strings.Split(strategy, " ")

		opp_move := move_map[strings.Trim(strs[0], " ")]
		your_move := move_map[strings.Trim(strs[1], " ")]

		// Calculate the score
		scores += score_map[your_move]

		// get the result vs opponent
		if your_move == "rock" && opp_move == "scissors" {
			scores += wl_map["win"]
		} else if your_move == "scissors" && opp_move == "paper" {
			scores += wl_map["win"]
		} else if your_move == "paper" && opp_move == "rock" {
			scores += wl_map["win"]
		} else if your_move == opp_move {
			scores += wl_map["draw"]
		} else {
			scores += wl_map["loss"]
		}

	}

	return scores
}

func RunDayTwoPartOne(filepath string) {
	strs := day1.ReadFileToArray(filepath)
	scores := PaperScissorsRock(strs)

	println("Day 2 Part 1: ", scores)
}

func PaperScissorsRockWinLoseDraw(strs []string) int {
	// Opponent's strategy
	// A -> Rock
	// B -> Paper
	// C -> Scissors
	// Your stragegy
	// X -> Rock
	// Y -> paper
	// z ->  scissors
	// Create a score map for every combination
	move_map := make(map[string]string)
	move_map["A"] = "rock"
	move_map["B"] = "paper"
	move_map["C"] = "scissors"

	// Result map
	result_map := make(map[string]string)
	result_map["X"] = "loss"
	result_map["Y"] = "draw"
	result_map["Z"] = "win"

	// create a score map for every move
	score_map := make(map[string]int)
	score_map["rock"] = 1
	score_map["paper"] = 2
	score_map["scissors"] = 3

	// Create a wl map fore winning and losing
	wl_map := make(map[string]int)
	wl_map["win"] = 6
	wl_map["draw"] = 3
	wl_map["loss"] = 0

	scores := 0
	for _, strategy := range strs {
		// Splite the line into your score and the opponent's score
		strs := strings.Split(strategy, " ")
		opp_move := move_map[strings.Trim(strs[0], " ")]
		desired_result := result_map[strings.Trim(strs[1], " ")]

		// do the move with the required result
		if desired_result == "win" {
			scores += wl_map["win"]
			if opp_move == "rock" {
				scores += score_map["paper"]
			} else if opp_move == "paper" {
				scores += score_map["scissors"]
			} else {
				scores += score_map["rock"]
			}
		} else if desired_result == "draw" {
			scores += wl_map["draw"]
			scores += score_map[opp_move]
		} else {
			scores += wl_map["loss"]
			if opp_move == "rock" {
				scores += score_map["scissors"]
			} else if opp_move == "paper" {
				scores += score_map["rock"]
			} else {
				scores += score_map["paper"]
			}
		}
	}

	return scores
}

func RunDayTwoPartTwo(filepath string) {
	strs := day1.ReadFileToArray(filepath)
	scores := PaperScissorsRockWinLoseDraw(strs)
	println("Day 2 Part 2: ", scores)
}
