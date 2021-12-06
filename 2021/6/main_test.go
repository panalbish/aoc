package main

import "testing"

func TestPart1(t *testing.T) {
	tt := []struct {
		fishTimerList []int
		totalDays     int
		expectedValue int
	}{
		{
			fishTimerList: []int{3, 4, 3, 1, 2},
			totalDays:     18,
			expectedValue: 26,
		},
		{
			fishTimerList: []int{3, 4, 3, 1, 2},
			totalDays:     80,
			expectedValue: 5934,
		},
		{
			fishTimerList: []int{3, 4, 3, 1, 2},
			totalDays:     256,
			expectedValue: 26984457539,
		},
	}

	for _, tc := range tt {
		t.Run("FishFarm", func(t *testing.T) {
			actual := simulateFishBreeding(tc.fishTimerList, tc.totalDays)
			if actual != tc.expectedValue {
				t.Errorf("Got %d, expected %d", actual, tc.expectedValue)
			}
		})
	}
}
