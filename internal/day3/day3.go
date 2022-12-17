package day3

import (
	"bufio"
	"io"
)

type (
	rucksack struct {
		items        []byte
		compartmant1 []byte
		compartmant2 []byte
	}
	rucksacks []rucksack
)

func Part1(in io.Reader) int {
	return parseInput(in).sumDiff()
}

func Part2(in io.Reader) int {
	return parseInput(in).sumBadges()
}

func parseInput(in io.Reader) rucksacks {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	result := make(rucksacks, 0)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		split := len(line) / 2
		r := rucksack{
			items:        line,
			compartmant1: line[:split],
			compartmant2: line[split:],
		}

		result = append(result, r)
	}

	return result
}

func (r rucksacks) sumDiff() int {
	sum := 0
	for _, sack := range r {
		sum += sack.sumDiff()
	}
	return sum
}

func (r rucksacks) sumBadges() int {
	sum := 0
	for i := 0; i < len(r); i += 3 {
		for _, item := range r[i].items {
			if contains(r[i+1].items, item) && contains(r[i+2].items, item) {
				sum += toNum(item)
				break
			}
		}
	}

	return sum
}

func (r rucksack) sumDiff() int {
	for _, b := range r.compartmant1 {
		if contains(r.compartmant2, b) {
			return toNum(b)
		}
	}

	return 0
}

func toNum(c byte) int {
	if c > 96 {
		return int(c - 96)
	}

	return int(c - 38)
}

func contains(compartment []byte, item byte) bool {
	for _, v := range compartment {
		if v == item {
			return true
		}
	}

	return false
}
