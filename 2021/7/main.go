package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1(positions []int) int {
	var result int
	sort.Ints(positions)
	medianFuelValue := medianNumber(positions)
	for _, position := range positions {
		result += diff(position, medianFuelValue)
	}
	return result
}

func Part2(positions []int) int {
	var result int
	sort.Ints(positions)
	averageFuelValue := averageNumber(positions)
	for _, position := range positions {
		for i := 1; i <= diff(position, averageFuelValue); i++ {
			result += i
		}
	}
	return result
}

func medianNumber(values []int) int {
	half := len(values) / 2
	if len(values)%2 == 0 {
		return (values[half-1] + values[half]) / 2
	}
	return values[half]
}

func averageNumber(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return int(math.Floor(float64(sum) / float64(len(values))))
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}

	inputLists := strings.Split(string(content), ",")
	var numberInputs []int
	for _, input := range inputLists {
		num, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}

		numberInputs = append(numberInputs, num)
	}

	fmt.Printf("Part 1 %v\n", Part1(numberInputs))
	fmt.Printf("Part 2 %v\n", Part2(numberInputs))
}
