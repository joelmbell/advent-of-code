package day11

import (
	"sort"
	"strconv"
	"strings"
)

type OP string

const (
	Add      OP = "+"
	Divide   OP = "/"
	Multiply OP = "*"
	Subtract OP = "-"
)

type operation struct {
	op    OP
	val   int
	isOld bool
}

func newOperation(opStr string) operation {
	op := OP(opStr[0:1])
	val, err := strconv.Atoi(opStr[2:])
	isOld := err != nil
	return operation{
		op, val, isOld,
	}
}

func (o operation) apply(toValue int) int {

	opValue := o.val
	if o.isOld {
		opValue = toValue
	}
	switch o.op {
	case Add:
		return toValue + opValue
	case Divide:
		return toValue / opValue
	case Multiply:
		return toValue * opValue
	case Subtract:
		return toValue - opValue
	default:
		panic("invalid op")
	}
}

type monkey struct {
	items          []int
	operation      operation
	testDivisor    int
	passToMonkey   int
	failToMonkey   int
	inspectedCount int
}

func parseMonkey(monkeyData []string) *monkey {
	monkey := monkey{}
	for _, item := range monkeyData {
		switch {
		case strings.Index(item, "Starting items: ") != -1:
			itemStr := item[18:]
			numStr := strings.Split(itemStr, ", ")
			items := make([]int, 0)
			for _, num := range numStr {
				conv, _ := strconv.Atoi(num)
				items = append(items, conv)
			}
			monkey.items = items
		case strings.Index(item, "Operation: ") != -1:
			operationStr := item[23:]
			operation := newOperation(operationStr)
			monkey.operation = operation
		case strings.Index(item, "Test: ") != -1:
			divisorStr := item[21:]
			divisor, _ := strconv.Atoi(divisorStr)
			monkey.testDivisor = divisor
		case strings.Index(item, "If true:") != -1:
			trueStr := item[len(item)-1:]
			to, _ := strconv.Atoi(trueStr)
			monkey.passToMonkey = to
		case strings.Index(item, "If false: ") != -1:
			falseStr := item[len(item)-1:]
			to, _ := strconv.Atoi(falseStr)
			monkey.failToMonkey = to
		default:
			continue
		}
	}

	return &monkey
}

func parse(input []string) []*monkey {
	monkeys := make([]*monkey, 0)
	cur := make([]string, 0)
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			monkeys = append(monkeys, parseMonkey(cur))
			cur = make([]string, 0)
		} else {
			cur = append(cur, input[i])
		}
	}
	monkeys = append(monkeys, parseMonkey(cur))
	return monkeys
}

func (m *monkey) inspect(shouldDivideBy3 bool) (throwTo int, worryLevel int) {
	toInspect := m.items[0]
	worryLevel = m.operation.apply(toInspect)
	if shouldDivideBy3 {
		worryLevel = int(float64(worryLevel) / float64(3))
	}
	if worryLevel%m.testDivisor == 0 {
		throwTo = m.passToMonkey
	} else {
		throwTo = m.failToMonkey
	}
	m.items = m.items[1:]
	m.inspectedCount += 1
	return
}

func Part1(input []string) int {
	monkeys := parse(input)

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			for len(monkey.items) > 0 {
				throwTo, worryLevel := monkey.inspect(true)
				monkeys[throwTo].items = append(monkeys[throwTo].items, worryLevel)
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedCount > monkeys[j].inspectedCount
	})

	return monkeys[0].inspectedCount * monkeys[1].inspectedCount
}

func Part2(input []string) int {
	monkeys := parse(input)

	for i := 0; i < 1000; i++ {
		for _, monkey := range monkeys {
			for len(monkey.items) > 0 {
				throwTo, worryLevel := monkey.inspect(false)
				monkeys[throwTo].items = append(monkeys[throwTo].items, worryLevel)
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedCount > monkeys[j].inspectedCount
	})

	return monkeys[0].inspectedCount * monkeys[1].inspectedCount
}
