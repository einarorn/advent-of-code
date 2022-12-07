package main

import (
	"fmt"
	"sort"
	"strconv"

	"advent-of-code/resource_provider"
)

type Snacks struct {
	Elf      int
	Calories int
}

func main() {
	input, _ := resource_provider.ReadAllLines("input.txt")

	var list []Snacks
	list = append(list, Snacks{Elf: 1, Calories: 0})
	i := 0

	for _, snack := range input {
		if snack != "" {
			calories, _ := strconv.Atoi(snack)
			list[i].Calories += calories
		} else {
			i++
			list = append(list, Snacks{Elf: i, Calories: 0})
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Calories > list[j].Calories
	})

	fmt.Println(list[0])
	fmt.Println(list[0].Calories + list[1].Calories + list[2].Calories)
}
