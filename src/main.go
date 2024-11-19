package main

import (
	"fmt"
	"go_day00/anscombe"
	"go_day00/cli"
)

func main() {
	var flags cli.MetricsFlags
	flags.Init()
	numbers, err := cli.GetNumbers()
	if err != nil {
		fmt.Println(err)
		return
	}
	if *flags.MeanPtr {
		fmt.Printf("Mean: %.2f\n", anscombe.Mean(numbers))
	}
	if *flags.MedianPtr {
		fmt.Printf("Median: %.2f\n", anscombe.Median(numbers))
	}
	if *flags.ModePtr {
		fmt.Printf("Moda: %d\n", anscombe.Mode(numbers))
	}
	if *flags.SDPtr {
		fmt.Printf("SD: %.2f\n", anscombe.StandardDeviation(numbers))
	}
}
