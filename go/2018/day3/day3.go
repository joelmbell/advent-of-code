package day3_2018

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Point models an x,y coordinate
type Point struct {
	X int
	Y int
}

func (p *Point) hash() string {
	xVal := strconv.Itoa(p.X)
	yVal := strconv.Itoa(p.Y)
	return xVal + "x" + yVal
}

func (p *Point) relativePoint(toPoint Point) Point {
	newX := p.X + toPoint.X
	newY := p.Y + toPoint.Y
	return Point{X: newX, Y: newY}
}

// Frame models the data to make a rectangle
type Frame struct {
	Point  Point
	Width  int
	Height int
}

// Value models the data at each point
type Value struct {
	ID    string
	Count int
}

// Claim models a pice of claim data
type Claim struct {
	ID    string
	Frame Frame
}

func (c *Claim) writeFrame() map[string]Value {
	store := make(map[string]Value)
	for y := 0; y < c.Frame.Height; y++ {
		for x := 0; x < c.Frame.Width; x++ {
			point := Point{X: x, Y: y}
			relativePoint := c.Frame.Point.relativePoint(point)
			store[relativePoint.hash()] = Value{ID: c.ID, Count: 1}
		}
	}

	return store
}

func createClaim(input string) Claim {
	atIndex := strings.Index(input, "@")
	commaIndex := strings.Index(input, ",")
	colonIndex := strings.Index(input, ":")
	xIndex := strings.Index(input, "x")
	endIndex := len(input)
	id := input[1 : atIndex-1]

	xCoord, _ := strconv.Atoi(input[atIndex+2 : commaIndex])
	yCoord, _ := strconv.Atoi(input[commaIndex+1 : colonIndex])
	width, _ := strconv.Atoi(input[colonIndex+2 : xIndex])
	height, _ := strconv.Atoi(input[xIndex+1 : endIndex])

	point := Point{X: xCoord, Y: yCoord}
	frame := Frame{Point: point, Width: width, Height: height}
	return Claim{ID: id, Frame: frame}
}

func loadData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []string
	for scanner.Scan() {
		string := scanner.Text()
		result = append(result, string)
	}

	return result
}

// Part1 solves part one with an arbitrary input
func Part1(input []string) int {
	store := make(map[string]Value)
	var result int

	for _, item := range input {
		claim := createClaim(item)
		claimStore := claim.writeFrame()
		for k, v := range claimStore {
			if _, ok := store[k]; ok {
				newCount := store[k].Count + 1
				store[k] = Value{ID: store[k].ID, Count: newCount}
				if newCount == 2 {
					result++
				}
			} else {
				store[k] = v
			}
		}
	}
	return result
}

// Part2 solves part 2
func Part2(input []string) string {

	store := make(map[string]Value)
	results := make(map[string]bool)

	for _, item := range input {
		claim := createClaim(item)
		var overlaps bool
		for k, v := range claim.writeFrame() {
			if _, ok := store[k]; ok {
				newCount := store[k].Count + 1
				store[k] = Value{ID: store[k].ID, Count: newCount}
				overlaps = true
			} else {
				store[k] = v
			}

			results[v.ID] = overlaps
			results[store[k].ID] = overlaps
		}
	}

	for k, v := range results {
		if v == false {
			return k
		}
	}

	return "unknown"
}

// Part1WithInput solves part 1
func Part1WithInput() int {
	data := loadData("input.txt")
	return Part1(data)
}

// Part2WithInput solves part 2
func Part2WithInput() string {
	data := loadData("input.txt")
	return Part2(data)
}
