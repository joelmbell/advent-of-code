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

type instruction struct {
	opponent string
	me       string
}

func (r *instruction) didWinPt1() bool {
	return r.me == "Y" && r.opponent == "A" || r.me == "Z" && r.opponent == "B" || r.me == "X" && r.opponent == "C"
}

func (r *instruction) didLosePt1() bool {
	return r.me == "Z" && r.opponent == "A" || r.me == "X" && r.opponent == "B" || r.me == "Y" && r.opponent == "C"
}

func (r *instruction) scorePt1() int {
	var score int
	switch r.me {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	}

	if r.didWinPt1() {
		score += 6
	} else if r.didLosePt1() {
		score += 0
	} else {
		score += 3
	}
	return score
}

func parse(input []string) []instruction {
	var rounds []instruction
	for _, item := range input {
		split := strings.Split(item, " ")
		rounds = append(rounds, instruction{
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
		total += round.scorePt1()
	}

	return total
}

func winsTo(input string) string {
	switch input {
	case "A":
		return "B"
	case "B":
		return "C"
	case "C":
		return "A"
	default:
		return ""
	}
}

func losesTo(input string) string {
	switch input {
	case "A":
		return "C"
	case "B":
		return "A"
	case "C":
		return "B"
	default:
		return ""
	}
}

func Part2(input []string) int {
	instructions := parse(input)

	var total int
	for _, instruction := range instructions {
		var play string
		switch instruction.me {
		case "X":
			play = losesTo(instruction.opponent)
			total += 0
		case "Y":
			play = instruction.opponent
			total += 3
		case "Z":
			play = winsTo(instruction.opponent)
			total += 6
		default:
			panic("invalid instruction")
		}

		//fmt.Printf("instruction: %v, \tplay: %v\n", instruction, play)

		switch play {
		case "A":
			total += 1
		case "B":
			total += 2
		case "C":
			total += 3
		}
	}
	return total
}
