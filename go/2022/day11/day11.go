package day11

import (
	"fmt"
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

func (m *monkey) inspect(shouldDivideBy3 bool, cycleLength int) (throwTo int, worryLevel int) {
	inspect := m.items[0]
	opapply := m.operation.apply(inspect)
	if shouldDivideBy3 {
		opapply = int(float64(opapply) / float64(3))
	}
	reduce := opapply % cycleLength
	worryLevel = reduce
	throwTo = m.passToMonkey
	if worryLevel%m.testDivisor != 0 {
		throwTo = m.failToMonkey
	}
	//fmt.Printf("\t inspect: %v, apply: %v, reduce: %v passTo: %v\n", inspect, opapply, reduce, throwTo)
	m.items = m.items[1:]
	m.inspectedCount += 1
	return
}

func Part1(input []string) int {
	monkeys := parse(input)
	cycleLength := calcCycleLength(monkeys)
	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			for len(monkey.items) > 0 {
				throwTo, worryLevel := monkey.inspect(true, cycleLength)
				monkeys[throwTo].items = append(monkeys[throwTo].items, worryLevel)
			}
		}
	}

	fmt.Printf("\n\n== RESULT ==\n")
	for idx, monkey := range monkeys {
		fmt.Printf("%v inspected: %v\n", idx, monkey.inspectedCount)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedCount > monkeys[j].inspectedCount
	})

	return monkeys[0].inspectedCount * monkeys[1].inspectedCount
}

func calcCycleLength(monkeys []*monkey) int {
	cycleLength := 1
	for _, monkey := range monkeys {
		cycleLength *= monkey.testDivisor
	}
	return cycleLength
}

func Part2(input []string) int {
	monkeys := parse(input)
	cycleLength := calcCycleLength(monkeys)

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			for len(monkey.items) > 0 {
				throwTo, worryLevel := monkey.inspect(false, cycleLength)
				monkeys[throwTo].items = append(monkeys[throwTo].items, worryLevel)
			}
		}
		// Print for debugging
		//round := i + 1
		//if round != 0 && (round == 1 || round == 20 || round%1000 == 0) {
		//	fmt.Printf("\n== After round %v ==\n\n", round)
		//	for idx, monkey := range monkeys {
		//		fmt.Printf("monkey %v inspected %v times\n", idx, monkey.inspectedCount)
		//	}
		//	fmt.Printf("\n")
		//}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedCount > monkeys[j].inspectedCount
	})

	return monkeys[0].inspectedCount * monkeys[1].inspectedCount
}
