package day9

import (
	ds "aoc/datastructures"
	"aoc/util"
	"fmt"
	"strconv"
)

type instruction struct {
	direction string
	count     int
}

func newInstruction(input string) instruction {
	dir := input[0:1]
	count, err := strconv.Atoi(input[2:])
	if err != nil {
		panic(err)
	}

	return instruction{
		direction: dir,
		count:     count,
	}
}

func parseInstructions(input []string) []instruction {
	instructions := make([]instruction, 0)
	for _, row := range input {
		instructions = append(instructions, newInstruction(row))
	}
	return instructions
}

type point struct {
	x int
	y int
}

func (p point) key() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

func (p point) isTouching(p2 point) bool {
	return util.Abs(p2.x-p.x) <= 1 && util.Abs(p2.y-p.y) <= 1
}

func moveDirection(p point, dir string) point {
	switch dir {
	case "R":
		p.x += 1
	case "U":
		p.y -= 1
	case "L":
		p.x -= 1
	case "D":
		p.y += 1
	default:
		panic(fmt.Sprintf("\tinvalid instruction %v", dir))
	}
	return p
}

func moveTail(tail point, head point) (newPoint point, didSkip bool) {

	if tail.isTouching(head) {
		return tail, true
	}

	ref := point{
		x: head.x - tail.x,
		y: head.y - tail.y,
	}

	if util.Abs(ref.x) == 2 {
		ref.x = util.TowardZero(ref.x)
	}

	if util.Abs(ref.y) == 2 {
		ref.y = util.TowardZero(ref.y)
	}

	tail = point{
		x: tail.x + ref.x,
		y: tail.y + ref.y,
	}
	return tail, false
}

var GridSize = point{x: 16, y: 16}

func draw(knots []point, maxKnots int) {

	knotsMap := make(map[string]string)
	for i, knot := range knots {
		knotName := ""
		switch i {
		case 0:
			knotName = "H"
			knotsMap[knot.key()] = "H"
			continue
		case maxKnots - 1:
			knotsMap[knot.key()] = "T"
		default:
			knotName = strconv.Itoa(i)
			knotsMap[knot.key()] = knotName
		}
	}

	for i := -(GridSize.y / 2); i < GridSize.y/2; i++ {
		fmt.Printf("\n%v\t", i)
		for k := -(GridSize.x / 2); k < GridSize.x/2; k++ {
			point := point{
				x: k, y: i,
			}
			if val, ok := knotsMap[point.key()]; ok {
				fmt.Printf(val)
			} else {
				fmt.Printf(".")
			}

		}
	}
	fmt.Printf("\n\n")
}

func Part1(input []string) int {
	instructions := parseInstructions(input)

	var head point
	var tail point

	tailTouched := ds.NewSet([]string{tail.key()})
	for i := 0; i < len(instructions); i++ {
		inst := instructions[i]
		for j := 0; j < inst.count; j++ {
			head = moveDirection(head, inst.direction)
			tail, _ = moveTail(tail, head)
			tailTouched = tailTouched.Add(tail.key())
		}
	}

	return len(tailTouched)
}

const MaxKnots = 10

func Part2(input []string) int {
	instructions := parseInstructions(input)

	knots := []point{
		{0, 0},
	}

	tailTouched := ds.NewSet([]string{knots[0].key()})
	for i := 0; i < len(instructions); i++ {
		inst := instructions[i]
		for j := 0; j < inst.count; j++ {
			for k := 0; k < MaxKnots; k++ {
				if k == 0 {
					knots[0] = moveDirection(knots[0], inst.direction)
					continue
				}
				if len(knots) <= k {
					knots = append(knots, point{x: 0, y: 0})
				}

				knots[k], _ = moveTail(knots[k], knots[k-1])

				if k == MaxKnots-1 {
					tailTouched = tailTouched.Add(knots[k].key())
				}
			}
		}
	}
	return len(tailTouched)
}
