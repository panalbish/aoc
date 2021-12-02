package main

import "testing"

func TestPart1(t *testing.T) {
	tt := []struct {
		values         []string
		expectedResult int
	}{
		{
			values:         nil,
			expectedResult: 0,
		},
		{
			values: []string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2"},
			expectedResult: 150,
		},
	}

	for _, tc := range tt {
		t.Run("Part1", func(t *testing.T) {
			if actual := Part1(tc.values); actual != tc.expectedResult {
				t.Fail()
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tt := []struct {
		values         []string
		expectedResult int
	}{
		{
			values:         nil,
			expectedResult: 0,
		},
		{
			values: []string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2"},
			expectedResult: 900,
		},
	}

	for _, tc := range tt {
		t.Run("Part2", func(t *testing.T) {
			if actual := Part2(tc.values); actual != tc.expectedResult {
				t.Fail()
			}
		})
	}
}
