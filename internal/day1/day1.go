package day1

import (
	"bufio"
	"io"
	"strconv"
)

func Part1(in io.Reader) int {
	return maxCalories(parseInput(in))
}

func Part2(in io.Reader) int {
	return sum(topN(parseInput(in), 3)...)
}

func parseInput(in io.Reader) []int {

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	calories := make([]int, 0)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			calories = append(calories, sum)
			sum = 0
			continue
		}

		cal, err := strconv.Atoi(line)
		sum += cal
		if err != nil {
			panic("Invalid input!!!")
		}
	}

	if sum > 0 {
		calories = append(calories, sum)
	}

	return calories
}

func maxCalories(calories []int) int {
	max := topN(calories, 1)
	if len(max) == 0 {
		return 0
	}
	return max[0]
}

func topN(calories []int, n int) []int {
	len := min(len(calories), n)

	top := make([]int, len)
	if len == 0 {
		return top
	}

	for _, cal := range calories {
		for i := range top {
			if cal > top[i] {
				top[i], cal = cal, top[i]
			}
		}
	}
	return top
}

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func sum(values ...int) (sum int) {
	for _, v := range values {
		sum += v
	}

	return
}
