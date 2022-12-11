package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"advent-of-code/resource_provider"
)

type Monkey struct {
	Items       []int
	Operator    string
	Operation   int
	Test        int
	IfTrue      int
	IfFalse     int
	Inspections int
}

func main() {
	input, _ := resource_provider.ReadAllLines("input.txt")

	monkeysPartOne := runInspections(parseMonkeyInput(input), 20, true)
	monkeysPartTwo := runInspections(parseMonkeyInput(input), 10_000, false)
	fmt.Printf("Part One solution: %d\n", calculateMonkeyBusinessLevel(monkeysPartOne))
	fmt.Printf("Part Two solution: %d\n", calculateMonkeyBusinessLevel(monkeysPartTwo))
}

func runInspections(monkeys []Monkey, rounds int, hasRelief bool) []Monkey {
	for round := 0; round < rounds; round++ {
		divider := 1
		for _, monkey := range monkeys {
			divider = lcm(divider, monkey.Test)
		}

		for i := 0; i < len(monkeys); i++ {
			calculateWorryLevel(monkeys[i], divider, hasRelief)
			for _, item := range monkeys[i].Items {
				monkeys[i].Inspections++
				receivingMonkey := monkeys[i].IfFalse
				if item%monkeys[i].Test == 0 {
					receivingMonkey = monkeys[i].IfTrue
				}
				monkeys[receivingMonkey].Items = append(monkeys[receivingMonkey].Items, item)
				monkeys[i].Items = monkeys[i].Items[1:]
			}
		}
	}
	return monkeys
}

func calculateWorryLevel(monkey Monkey, divider int, hasRelief bool) {
	for i, item := range monkey.Items {
		operation := monkey.Operation
		if operation == -1 {
			operation = item
		}

		if monkey.Operator == "*" {
			monkey.Items[i] = item * operation
		} else if monkey.Operator == "+" {
			monkey.Items[i] = item + operation
		}

		if hasRelief {
			monkey.Items[i] = monkey.Items[i] / 3
		} else {
			monkey.Items[i] = monkey.Items[i] % divider
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func calculateMonkeyBusinessLevel(monkeys []Monkey) int {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspections > monkeys[j].Inspections
	})
	return monkeys[0].Inspections * monkeys[1].Inspections
}

func stringItemsToInt(items []string) []int {
	if len(items) == 0 {
		return []int{-1}
	}

	var intItems []int
	for _, item := range items {
		value, _ := strconv.Atoi(item)
		intItems = append(intItems, int(value))
	}
	return intItems
}

func parseMonkeyInput(input []string) []Monkey {
	var monkeyBusiness []Monkey
	reOperator := regexp.MustCompile(`\*|\+`)
	reNum := regexp.MustCompile(`[0-9]+`)

	for i := 0; i < len(input); i = i + 7 {
		monkey := Monkey{}
		monkey.Items = stringItemsToInt(reNum.FindAllString(input[i+1], -1))
		monkey.Operator = reOperator.FindAllString(input[i+2], -1)[0]
		monkey.Operation = stringItemsToInt(reNum.FindAllString(input[i+2], -1))[0]
		monkey.Test = stringItemsToInt(reNum.FindAllString(input[i+3], -1))[0]
		monkey.IfTrue = stringItemsToInt(reNum.FindAllString(input[i+4], -1))[0]
		monkey.IfFalse = stringItemsToInt(reNum.FindAllString(input[i+5], -1))[0]
		monkeyBusiness = append(monkeyBusiness, monkey)
	}
	return monkeyBusiness
}