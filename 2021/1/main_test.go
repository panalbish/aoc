package main

import "testing"

func TestPart1(t *testing.T) {
	tt := []struct {
		values         []int
		expectedResult int
	}{
		{
			values:         nil,
			expectedResult: 0,
		},
		{
			values:         []int{10, 4, 13},
			expectedResult: 1,
		},
		{
			values:         []int{1, 4, 13, 2, 20, 15, 200, 250},
			expectedResult: 5,
		},
		{
			values:         []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			expectedResult: 7,
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
		values         []int
		expectedResult int
	}{
		{
			values:         nil,
			expectedResult: 0,
		},
		{
			values:         []int{1, 5, 4, 3, 7, 10},
			expectedResult: 3,
		},
		{
			values:         []int{607, 618, 618, 617, 647, 716, 769, 792},
			expectedResult: 5,
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
