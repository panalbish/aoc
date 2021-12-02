package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Part1(values []string) int {
	position := map[string]int{
		"x": 0,
		"y": 0,
	}
	for _, value := range values {
		if value == "" {
			continue
		}

		split := strings.Split(value, " ")
		direction := split[0]
		unit := split[1]
		num, err := strconv.Atoi(unit)
		if err != nil {
			continue
		}

		switch direction {
		case "forward":
			position["x"] += num
		case "up":
			position["y"] -= num
		case "down":
			position["y"] += num
		}
	}
	return position["x"] * position["y"]
}

func Part2(values []string) int {
	position := map[string]int{
		"x":   0,
		"y":   0,
		"aim": 0,
	}
	for _, value := range values {
		if value == "" {
			continue
		}

		split := strings.Split(value, " ")
		direction := split[0]
		unit := split[1]
		num, err := strconv.Atoi(unit)
		if err != nil {
			continue
		}

		switch direction {
		case "forward":
			position["x"] += num
			if position["aim"] != 0 {
				position["y"] += position["aim"] * num
			}
		case "up":
			position["aim"] -= num
		case "down":
			position["aim"] += num
		}
	}
	return position["x"] * position["y"]
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
	fmt.Println(Part1(inputLists))
	fmt.Println(Part2(inputLists))
}
