package cli

import (
	"bufio"
	"errors"
	"flag"
	"os"
	"sort"
	"strconv"
)

func ValildNumber(n int) bool {
	if n < -100000 || n > 100000 {
		return false
	}
	return true
}

func GetNumbers() ([]int, error) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	numbers := make([]int, 0)
	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			return nil, err
		}
		if !ValildNumber(num) {
			return nil, errors.New("invalid number")
		}
		numbers = append(numbers, num)
	}
	if len(numbers) == 0 {
		return nil, errors.New("no numbers")
	}
	sort.Ints(numbers)
	return numbers, nil
}

type MetricsFlags struct {
	MeanPtr   *bool
	ModePtr   *bool
	MedianPtr *bool
	SDPtr     *bool
}

func (b *MetricsFlags) Init() {
	b.MeanPtr = flag.Bool("mean", false, "Find mean")
	b.ModePtr = flag.Bool("mode", false, "Find moda")
	b.MedianPtr = flag.Bool("median", false, "Find median")
	b.SDPtr = flag.Bool("sd", false, "Find standard deviation")
	line := os.Args[1:]
	if len(line) == 0 {
		*b.MeanPtr = true
		*b.ModePtr = true
		*b.MedianPtr = true
		*b.SDPtr = true
	}
	flag.Parse()
}
