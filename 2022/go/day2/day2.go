package day2

import (
	ds "2022/datastructures"
	"strings"
)

type instruction struct {
	opponent string
	me       string
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

var pointsForPlay = map[string]int{
	"A": 1, "B": 2, "C": 3,
}

func Part1(input []string) int {
	plays := ds.NewCircularLinkedList("A", "B", "C")
	playsMap := map[string]string{
		"X": "A", "Y": "B", "Z": "C",
	}

	var total int
	for _, round := range parse(input) {
		me := plays.Find(playsMap[round.me])
		opponent := plays.Find(round.opponent)

		switch {
		case me.Next.Value == opponent.Value:
			total += 0
		case me.Value == opponent.Value:
			total += 3
		case me.Previous.Value == opponent.Value:
			total += 6
		}

		total += pointsForPlay[me.Value]
	}

	return total
}

func Part2(input []string) int {
	plays := ds.NewCircularLinkedList("A", "B", "C")
	instructions := parse(input)

	var total int
	for _, instruction := range instructions {
		var play string
		switch instruction.me {
		case "X":
			play = plays.Find(instruction.opponent).Previous.Value
			total += 0
		case "Y":
			play = instruction.opponent
			total += 3
		case "Z":
			play = plays.Find(instruction.opponent).Next.Value
			total += 6
		default:
			panic("invalid instruction")
		}
		total += pointsForPlay[play]
	}
	
	return total
}
