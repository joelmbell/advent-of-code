package day3

import (
	"bufio"
	// "fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

type Frame struct {
	Point  Point
	Width  int
	Height int
}

func (f *Frame) writeFrame() map[string]int {
	store := make(map[string]int)
	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			point := Point{X: x, Y: y}
			relativePoint := f.Point.relativePoint(point)
			store[relativePoint.hash()] = 1
		}
	}

	return store
}

type Claim struct {
	ID    string
	Frame Frame
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

func Part1(input []string) int {

	store := make(map[string]int)
	for _, item := range input {
		claim := createClaim(item)
		claimStore := claim.Frame.writeFrame()
		for k, _ := range claimStore {
			if _, ok := store[k]; ok {
				store[k] = store[k] + 1
			} else {
				store[k] = 1
			}
		}
	}

	var result int
	for k, _ := range store {
		if store[k] > 1 {
			result++
		}
	}

	return result
}

func Part1WithInput() int {
	data := loadData("input.txt")
	return Part1(data)
}
