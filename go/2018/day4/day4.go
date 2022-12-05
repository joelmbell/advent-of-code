package day4_2018

import (
	"fmt"
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
	GuardID int // -1 if we don't know yet
	Type    EventType
}

type GuardData struct {
	TotalAsleep  time.Duration
	PerMinAsleep map[int]int
}

func createEvent(s string) Event {

	closeBracketIndex := strings.Index(s, "]")
	var eventType EventType
	guardID := -1
	switch {
	case strings.Index(s, "Guard") != -1:
		numCharIndex := strings.Index(s, "#")
		beginsShiftIndex := strings.Index(s, "begins shift")
		guardID, _ = strconv.Atoi(s[numCharIndex+1 : beginsShiftIndex-1])
		eventType = BeginShift
	case strings.Index(s, "wakes up") != -1:
		eventType = WakesUp
		guardID = -1
	case strings.Index(s, "falls asleep") != -1:
		eventType = FallsAsleep
		guardID = -1
	}

	dateString := s[1:closeBracketIndex]
	time, _ := time.Parse("2006-01-02 15:04", dateString)
	return Event{Time: time, GuardID: guardID, Type: eventType}
}

func parse(input []string) []Event {
	var events []Event
	for _, item := range input {
		e := createEvent(item)
		events = append(events, e)
	}
	return events
}

func groupDataByGuards(input []string) map[int]GuardData {
	events := parse(input)
	sort.Slice(events, func(i int, j int) bool {
		return events[i].Time.Before(events[j].Time)
	})

	guardGroup := make(map[int][]Event)
	var current int
	for _, e := range events {
		switch e.Type {
		case BeginShift:
			if _, ok := guardGroup[e.GuardID]; !ok {
				guardGroup[e.GuardID] = make([]Event, 0)
			}
			current = e.GuardID
		case FallsAsleep, WakesUp:
			e.GuardID = current
			guardGroup[current] = append(guardGroup[current], e)
		}
	}

	guardData := make(map[int]GuardData)
	for guardId, events := range guardGroup {
		var startSleep time.Time
		var durationAsleep time.Duration
		perMinAsleep := make(map[int]int)
		for _, e := range events {
			switch e.Type {
			case FallsAsleep:
				startSleep = e.Time
			case WakesUp:
				durationAsleep += e.Time.Sub(startSleep)
				for f := startSleep; e.Time.After(f); f = f.Add(60000000000) {
					if _, ok := perMinAsleep[f.Minute()]; !ok {
						perMinAsleep[f.Minute()] = 1
					} else {
						perMinAsleep[f.Minute()] += 1
					}
				}
			}
		}
		guardData[guardId] = GuardData{
			TotalAsleep:  durationAsleep,
			PerMinAsleep: perMinAsleep,
		}
	}
	return guardData
}

func Part1(input []string) int {
	guardData := groupDataByGuards(input)

	var currentID int
	var currentMax time.Duration
	for guardId, data := range guardData {
		if data.TotalAsleep > currentMax {
			currentID = guardId
			currentMax = data.TotalAsleep
		}
	}

	var curHighestMin int
	var curHighestCount int
	for k, v := range guardData[currentID].PerMinAsleep {
		if v > curHighestCount {
			curHighestMin = k
			curHighestCount = v
		}
	}
	return currentID * curHighestMin
}

func Part2(input []string) int {
	guardData := groupDataByGuards(input)

	maxGuard := -1
	maxMin := 0
	maxAmountSlept := 0
	for guardId, data := range guardData {
		for min, amountSlept := range data.PerMinAsleep {
			if amountSlept > maxAmountSlept {
				maxGuard = guardId
				maxMin = min
				maxAmountSlept = amountSlept
			}
		}
	}
	fmt.Printf("maxGuard: %v, maxMin: %v", maxGuard, maxMin)
	return maxGuard * maxMin
}
