package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Part1(values []int) int {
	if len(values) == 0 {
		return 0
	}

	prevVal := values[0]
	result := 0
	for _, val := range values {
		if val > prevVal {
			result++
		}
		prevVal = val
	}
	return result
}

func Part2(values []int) int {
	if len(values) == 0 {
		return 0
	}

	measurement := 3
	result := 0

	slidingWindows := make([][]int, 0)
	max := len(values) - measurement + 1
	for key := range values[:max] {
		slidingWindows = append(slidingWindows, values[key:key+measurement])
	}

	prevWindowSum := sum(slidingWindows[0])

	for _, value := range slidingWindows[1:] {
		totalValue := sum(value)
		if totalValue > prevWindowSum {
			result++
		}
		prevWindowSum = totalValue
	}

	return result
}

func sum(values []int) int {
	result := 0
	for _, value := range values {
		result += value
	}
	return result
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

	inputLists := strings.Split(string(content), "\n")
	var numberInputs []int
	for _, input := range inputLists {
		num, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		numberInputs = append(numberInputs, num)
	}

	fmt.Println(Part1(numberInputs))
	fmt.Println(Part2(numberInputs))
}
