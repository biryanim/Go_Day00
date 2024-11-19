package anscombe

import (
	"math"
	"sort"
	"testing"
)

func equalFloat(a, b float64) bool {
	return math.Abs(a-b) < 0.01
}

func TestMetrics(t *testing.T) {
	tests := []struct {
		numbers     []int
		rightMean   float64
		rightMedian float64
		rightMode   int
		rightSD     float64
	}{
		{[]int{10, 2, 38, 23, 38, 23, 21}, 22.14, 23.00, 23, 12.30},
		{[]int{10, 12, 23, 23, 16, 23, 21, 16}, 18.00, 18.50, 23, 4.90},
		{[]int{-2, 3, -10, 3, 100, 100, -10, 68, 44}, 32.89, 3.00, -10, 43.52},
		{[]int{68, -76, 31, -98, -95, 24, -41, -42, 12, -79}, -29.6, -41.5, -98, 56.24},
		{[]int{-47, -85, 40, -45, -28, 57, 39, -22, -41, 39}, -9.3, -25, 39, 46.30},
		{[]int{-37, 96, -35, 74, -76, 20, 74, 1, 31, 91, -45, 8, -81, 48, -81}, 5.87, 8.0, -81, 60.65},
		{[]int{65421, 7946, 59612, -67852, -35318, -44663, 44165, 38043, 65714, -71825, -71266, 92942, -73047, 22247, -72126}, -2667.13, 7946.0, -73047, 59588.01},
		{[]int{-20, -10, -18, -19, 2, 13, 13, 4, -5, 3, -13, -10, 3, 5, 12}, -2.67, 2, -10, 11.31},
		{[]int{-20, 15, -19, -16, -4, -8, 19, -7, -5, 17, -6, -18, 19, 7, -3, 7, 5, 3, -11, 9}, -0.8, -3.5, 7, 12.44},
	}

	for _, test := range tests {
		sort.Ints(test.numbers)
		mean := Mean(test.numbers)
		median := Median(test.numbers)
		mode := Mode(test.numbers)
		sd := StandardDeviation(test.numbers)
		if !equalFloat(mean, test.rightMean) {
			t.Errorf("Test Mean: %f != %f\n", mean, test.rightMean)
		}
		if !equalFloat(median, test.rightMedian) {
			t.Errorf("Test Median: %f != %f\n", median, test.rightMedian)
		}
		if mode != test.rightMode {
			t.Errorf("Test Mode: %d != %d\n", mode, test.rightMode)
		}
		if !equalFloat(sd, test.rightSD) {
			t.Errorf("Test Standard Deviation: %f != %f\n", sd, test.rightSD)
		}
	}
}
