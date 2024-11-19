package anscombe

import "math"

func Mean(nums []int) float64 {
	var avg float64
	for _, num := range nums {
		avg += float64(num)
	}
	return avg / float64(len(nums))
}

func Median(nums []int) float64 {
	if len(nums)%2 == 0 {
		return (float64(nums[len(nums)/2]) + float64(nums[len(nums)/2-1])) / 2.0
	} else {
		return float64(nums[len(nums)/2])
	}
}

func Mode(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	ans, cnt := nums[0], m[nums[0]]
	for k, value := range m {
		if cnt < value {
			ans = k
			cnt = value
		} else if cnt == value && k < ans {
			ans = k
		}
	}
	return ans
}

func StandardDeviation(nums []int) float64 {
	var res float64
	mean := Mean(nums)
	for _, num := range nums {
		res += math.Pow(float64(num)-mean, 2)
	}
	return math.Sqrt(res / float64(len(nums)))
}
