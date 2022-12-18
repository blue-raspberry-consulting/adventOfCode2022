package day4

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type (
	assignment struct {
		start int
		end   int
	}

	pair [2]assignment

	pairs []pair
)

func Part1(in io.Reader) int {
	return parseInput(in).countFullContained()
}

func Part2(in io.Reader) int {
	return parseInput(in).countOverlap()
}

func parseInput(in io.Reader) pairs {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	result := make(pairs, 0)
	for scanner.Scan() {
		line := scanner.Text()
		assignments := strings.Split(line, ",")

		result = append(result, pair{
			parseAssignments(assignments[0]),
			parseAssignments(assignments[1]),
		})
	}

	return result
}

func parseAssignments(ass string) assignment {
	v := strings.Split(ass, "-")
	start, _ := strconv.Atoi(v[0])
	end, _ := strconv.Atoi(v[1])
	return assignment{start: start, end: end}
}

func (p pairs) countFullContained() int {
	count := 0
	for _, v := range p {
		if v.isFullContained() {
			count++
		}
	}
	return count
}

func (p pairs) countOverlap() int {
	count := 0
	for _, v := range p {
		if v.isOverlap() {
			count++
		}
	}
	return count
}

func (p pair) isFullContained() bool {

	if (p[0].start <= p[1].start && p[0].end >= p[1].end) || (p[1].start <= p[0].start && p[1].end >= p[0].end) {
		return true
	}

	return false
}

func (p pair) isOverlap() bool {
	if (p[1].start >= p[0].start && p[1].start <= p[0].end) || (p[0].start >= p[1].start && p[0].start <= p[1].end) {
		return true
	}

	return false
}
