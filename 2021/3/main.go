package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Part1(values []string) int64 {
	var gammaResult string
	var epsilonResult string

	for key := range values[0] {
		var countOfOnes int
		var countOfZeroes int
		for _, value := range values {
			switch string(value[key]) {
			case "0":
				countOfZeroes++
			case "1":
				countOfOnes++
			}
		}
		if countOfOnes > countOfZeroes {
			gammaResult += "1"
			epsilonResult += "0"
		} else {
			gammaResult += "0"
			epsilonResult += "1"
		}
	}

	gammaDecimal, err := strconv.ParseInt(gammaResult, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilonDecimal, err := strconv.ParseInt(epsilonResult, 2, 64)
	if err != nil {
		panic(err)
	}

	return gammaDecimal * epsilonDecimal
}

type ratingData struct {
	allValues  []string
	oneValues  []string
	zeroValues []string
	index      int
}

func oxygenGeneratorRating(data *ratingData) {
	data.oneValues = nil
	data.zeroValues = nil

	for _, value := range data.allValues {
		switch string(value[data.index]) {
		case "0":
			data.zeroValues = append(data.zeroValues, value)
		case "1":
			data.oneValues = append(data.oneValues, value)
		}
	}

	if len(data.allValues) == 1 {
		return
	}

	if len(data.oneValues) > len(data.zeroValues) {
		data.allValues = data.oneValues
	}

	if len(data.oneValues) < len(data.zeroValues) {
		data.allValues = data.zeroValues
	}

	if len(data.oneValues) == len(data.zeroValues) {
		data.allValues = data.oneValues
	}

	return
}

func co2GeneratorRating(data *ratingData) {
	data.oneValues = nil
	data.zeroValues = nil

	for _, value := range data.allValues {
		switch string(value[data.index]) {
		case "0":
			data.zeroValues = append(data.zeroValues, value)
		case "1":
			data.oneValues = append(data.oneValues, value)
		}
	}

	if len(data.allValues) == 1 {
		return
	}

	if len(data.oneValues) > len(data.zeroValues) {
		data.allValues = data.zeroValues
	}

	if len(data.oneValues) < len(data.zeroValues) {
		data.allValues = data.oneValues
	}

	if len(data.oneValues) == len(data.zeroValues) {
		data.allValues = data.zeroValues
	}
}

func Part2(values []string) int64 {
	oxygenOperation := ratingData{
		allValues: values,
	}

	co2Operation := ratingData{
		allValues: values,
	}

	for key := range values[0] {
		oxygenOperation.index = key
		co2Operation.index = key

		oxygenGeneratorRating(&oxygenOperation)
		co2GeneratorRating(&co2Operation)
	}

	oxygenResultDecimal, err := strconv.ParseInt(oxygenOperation.allValues[0], 2, 64)
	if err != nil {
		panic(err)
	}

	co2ResultDecimal, err := strconv.ParseInt(co2Operation.allValues[0], 2, 64)
	if err != nil {
		panic(err)
	}

	return oxygenResultDecimal * co2ResultDecimal
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
