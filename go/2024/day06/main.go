package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	"github.com/YourAverageMoron/aoc/go/lib/app"
)

func main() {
	logger := slog.New(slog.Default().Handler())
	app, err := app.New(logger)
	if err != nil {
		logger.Error("error initialising app", slog.Any("error", err))
		return
	}
	err = app.Run(guardGallivant)
	if err != nil {
		logger.Error(err.Error())
	}
	err = app.Run(guardGallivant2)
	if err != nil {
		logger.Error(err.Error())
	}
}

func guardGallivant(f *os.File) (string, error) {
	m := parseFile(f)

	i, j := findGuard(m)

	seen := traverseMap(i, j, m, make(map[string]bool))

	sum := 0
	for range seen {
		sum++
	}
	return fmt.Sprint(sum), nil
}

func guardGallivant2(f *os.File) (string, error) {
	m := parseFile(f)
	i, j := findGuard(m)

	loops := findLoops(i, j, m, make(map[string]bool))
	fmt.Println(loops)

	return "", nil
}

func findLoops(currI, currJ int, m [][]rune, seen map[string]bool) int {
	seen[fmt.Sprintf("%d,%d", currI, currJ)] = true
	currDir := m[currI][currJ]
	nextI, nextJ := getNextLocation(currI, currJ, currDir)

	// TODO: this may need to check that the next location is not out of bounds
	if !seen[fmt.Sprintf("%d,%d", nextI, nextJ)] && m[nextI][nextJ] != '#' {
        
	}
	// If not in seen -> add block to next location
	// Run loop check
	// If loop check add to loops
	return 1
}

func getNextLocation(i, j int, d rune) (int, int) {
	switch d {
	case '^':
		return i - 1, j
	case '>':
		return i, j + 1
	case 'v':
		return i + 1, j
	case '<':
		return i, j - 1
	default:
		return -1, -1
	}
}
func traverseMap(currI, currJ int, m [][]rune, seen map[string]bool) map[string]bool {
	seen[fmt.Sprintf("%d,%d", currI, currJ)] = true

	currDir := m[currI][currJ]

	switch currDir {
	case '^':
		if currI == 0 {
			return seen
		}
		if m[currI-1][currJ] == '#' {
			m[currI][currJ] = '>'
			return traverseMap(currI, currJ, m, seen)
		}
		m[currI][currJ] = '.'
		m[currI-1][currJ] = currDir
		return traverseMap(currI-1, currJ, m, seen)
	case '>':
		if currJ == len(m[currI])-1 {
			return seen
		}
		if m[currI][currJ+1] == '#' {
			m[currI][currJ] = 'v'
			return traverseMap(currI, currJ, m, seen)
		}
		m[currI][currJ] = '.'
		m[currI][currJ+1] = currDir
		return traverseMap(currI, currJ+1, m, seen)
	case 'v':
		if currI == len(m)-1 {
			return seen
		}
		if m[currI+1][currJ] == '#' {
			m[currI][currJ] = '<'
			return traverseMap(currI, currJ, m, seen)
		}
		m[currI][currJ] = '.'
		m[currI+1][currJ] = currDir
		return traverseMap(currI+1, currJ, m, seen)
	case '<':
		if currJ == 0 {
			return seen
		}
		if m[currI][currJ-1] == '#' {
			m[currI][currJ] = '^'
			return traverseMap(currI, currJ, m, seen)
		}
		m[currI][currJ] = '.'
		m[currI][currJ-1] = currDir
		return traverseMap(currI, currJ-1, m, seen)
	default:
		return seen
	}
}

func findGuard(m [][]rune) (int, int) {
	for i, row := range m {
		for j, val := range row {
			if val == '^' || val == '>' || val == 'v' || val == '<' {
				return i, j
			}
		}
	}
	return -1, -1
}

func parseFile(f *os.File) [][]rune {
	scanner := bufio.NewScanner(f)
	var res [][]rune

	i := 0
	for scanner.Scan() {
		res = append(res, make([]rune, 0))
		for _, val := range scanner.Text() {
			res[i] = append(res[i], val)
		}
		i++
	}
	return res
}
