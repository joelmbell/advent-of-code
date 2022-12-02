package day2

import (
	"strings"
)

/*
	first colum
	A - Rock
	B - Paper
	C - Scissors

	second column
	X - Rock
	Y - Paper
	Z - Scissors

	total score
	- sum of your scores for each round
		1 - rock
		2 - paper
		3 - scissors
	- Plus outcome of the round
		0 - lost
		3 - draw
		6 - won
*/

type round struct {
	opponent string
	me       string
}

func (r *round) didWin() bool {
	return r.me == "Y" && r.opponent == "A" || r.me == "Z" && r.opponent == "B" || r.me == "X" && r.opponent == "C"
}

func (r *round) didLose() bool {
	return r.me == "Z" && r.opponent == "A" || r.me == "X" && r.opponent == "B" || r.me == "Y" && r.opponent == "C"
}

func (r *round) score() int {
	var score int
	switch r.me {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	}

	if r.didWin() {
		score += 6
	} else if r.didLose() {
		score += 0
	} else {
		score += 3
	}
	return score
}

func parse(input []string) []round {
	var rounds []round
	for _, item := range input {
		split := strings.Split(item, " ")
		rounds = append(rounds, round{
			opponent: split[0],
			me:       split[1],
		})
	}
	return rounds
}

func Part1(input []string) int {
	rounds := parse(input)

	var total int
	for _, round := range rounds {
		total += round.score()
	}

	return total
}
