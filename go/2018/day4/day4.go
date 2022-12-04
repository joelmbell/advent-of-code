package day4_2018

import (
	"bufio"
	// "fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type EventType string

const (
	BeginShift  EventType = "begin"
	FallsAsleep EventType = "asleep"
	WakesUp     EventType = "wake"
)

type Event struct {
	Time    time.Time
	GuardID int // -1 if no guardID is available
	Type    EventType
}

type GuardData struct {
	// ID           int
	TotalAsleep  time.Duration
	PerMinAsleep map[int]time.Duration
}

func createEvent(s string) Event {

	closeBracketIndex := strings.Index(s, "]")
	guardID := -1
	var eventType EventType

	switch {
	case strings.Index(s, "Guard") != -1:
		numCharIndex := strings.Index(s, "#")
		beginsShiftIndex := strings.Index(s, "begins shift")
		guardID, _ = strconv.Atoi(s[numCharIndex+1 : beginsShiftIndex-1])
		eventType = BeginShift
	case strings.Index(s, "wakes up") != -1:
		eventType = WakesUp
	case strings.Index(s, "falls asleep") != -1:
		eventType = FallsAsleep
	}

	dateString := s[1:closeBracketIndex]
	time, _ := time.Parse("2006-01-02 15:04", dateString)
	return Event{Time: time, GuardID: guardID, Type: eventType}
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
	var events []Event
	for _, item := range input {
		e := createEvent(item)
		events = append(events, e)
	}

	sort.Slice(events, func(i int, j int) bool {
		return events[i].Time.Before(events[j].Time)
	})

	// key = guardID
	store := make(map[int]GuardData)

	currentGuard := -1
	var fellAsleep time.Time
	for _, e := range events {
		switch e.Type {
		case BeginShift:
			currentGuard = e.GuardID
		case FallsAsleep:
			fellAsleep = e.Time
		case WakesUp:
			durationAsleep := e.Time.Sub(fellAsleep)
			data := store[currentGuard]
			newTotalAsleep := data.TotalAsleep + durationAsleep
			newData := GuardData{TotalAsleep: newTotalAsleep}

			// if day is the same as when you feel asleep
			// write all the minutes he was asleep
			// if day is different,
			// start at beginning of today, and write all the min he was asleep

			store[currentGuard] = newData
		}
	}

	var currentID int
	var currentMax time.Duration
	for k, v := range store {
		if v.TotalAsleep > currentMax {
			currentID = k
			currentMax = v.TotalAsleep
		}
	}

	return currentID
}

func Part1WithInput(filename string) int {
	data := loadData(filename)
	return Part1(data)
}
