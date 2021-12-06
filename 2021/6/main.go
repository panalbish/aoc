package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func simulateFishBreeding(fishTimerList []int, totalDays int) int {
	var result int
	fishAgeGroups := [9]int{}
	for _, fishTimer := range fishTimerList {
		fishAgeGroups[fishTimer]++
	}

	for i := 0; i < totalDays; i++ {
		bredFish := fishAgeGroups[0]
		for key := range fishAgeGroups[:len(fishAgeGroups)-1] {
			fishAgeGroups[key] = fishAgeGroups[key+1]
		}
		fishAgeGroups[6] += bredFish
		fishAgeGroups[8] = bredFish
	}

	for _, fish := range fishAgeGroups {
		result += fish
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

	inputLists := strings.Split(string(content), ",")
	var numberInputs []int
	for _, input := range inputLists {
		num, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}

		numberInputs = append(numberInputs, num)
	}

	fmt.Printf("Part 1 %v\n", simulateFishBreeding(numberInputs, 80))

	fmt.Printf("Part 2 %v\n", simulateFishBreeding(numberInputs, 256))
}
