package day2

import (
	"bufio"
	"io"
	"strings"
)

var (
	rock     = weapon{score: 1}
	paper    = weapon{score: 2}
	scissors = weapon{score: 3}

	win   = outcome{score: 6}
	draw  = outcome{score: 3}
	loose = outcome{score: 0}

	weapons = map[string]weapon{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	loosesTo = map[weapon]weapon{
		rock:     paper,
		paper:    scissors,
		scissors: rock,
	}

	winsTo = map[weapon]weapon{
		rock:     scissors,
		paper:    rock,
		scissors: paper,
	}

	outcomes = map[string]outcome{
		"X": loose,
		"Y": draw,
		"Z": win,
	}
)

type (
	weapon struct {
		score int
	}
	outcome struct {
		score int
	}
	game struct {
		elfWeapon     weapon
		myWeapon      weapon
		wantedOutcome outcome
	}
	games []game
)

func Part1(in io.Reader) int {
	return parseInput(in).getScore(false)
}

func Part2(in io.Reader) int {
	return parseInput(in).getScore(true)
}

func parseInput(in io.Reader) games {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	result := make(games, 0)
	for scanner.Scan() {
		line := scanner.Text()
		w := strings.Split(line, " ")
		g := game{
			elfWeapon:     weapons[w[0]],
			myWeapon:      weapons[w[1]],
			wantedOutcome: outcomes[w[1]],
		}

		result = append(result, g)
	}

	return result
}

func (g games) getScore(useOutcome bool) int {
	score := 0
	for _, game := range g {
		if useOutcome {
			score += game.getScore2()
		} else {
			score += game.getScore()
		}

	}

	return score
}

func (g game) getScore() int {
	score := g.myWeapon.score

	if g.myWeapon == g.elfWeapon {
		score += 3
	} else if loosesTo[g.elfWeapon] == g.myWeapon {
		score += 6
	}

	return score
}

func (g game) getScore2() int {
	score := g.wantedOutcome.score

	if g.wantedOutcome == draw {
		score += g.elfWeapon.score
	} else if g.wantedOutcome == win {
		score += loosesTo[g.elfWeapon].score
	} else {
		score += winsTo[g.elfWeapon].score
	}

	return score
}
