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
	operation  string
	index      int
}

func oxygenGeneratorRating(data *ratingData) ratingData {
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
		return *data
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

	return *data
}

func co2GeneratorRating(data *ratingData) ratingData {
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
		return *data
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

	return *data
}

func Part2(values []string) int64 {
	var oxygenResult ratingData
	var co2Result ratingData

	oxygenOperation := ratingData{
		operation: "o2",
		allValues: values,
	}

	co2Operation := ratingData{
		operation: "co2",
		allValues: values,
	}

	for key := range values[0] {
		oxygenOperation.index = key
		co2Operation.index = key

		oxygenResult = oxygenGeneratorRating(&oxygenOperation)
		co2Result = co2GeneratorRating(&co2Operation)
	}

	oxygenResultDecimal, err := strconv.ParseInt(oxygenResult.allValues[0], 2, 64)
	if err != nil {
		panic(err)
	}

	co2ResultDecimal, err := strconv.ParseInt(co2Result.allValues[0], 2, 64)
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
