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
	err = app.Run(ceresSearch)
	if err != nil {
		logger.Error(err.Error())
	}
	err = app.Run(ceresSearch2)
	if err != nil {
		logger.Error(err.Error())
	}
}

func ceresSearch2(f *os.File) (string, error) {
	file := getFile(f)

	sum := 0
	for i, line := range file {
		for j, char := range line {
			if char == rune('A') {
				sum += findMasX(file, i, j)
			}
		}
	}
	return fmt.Sprint(sum), nil
}

func findMasX(file [][]rune, aI, aJ int) int {
	// Left up
	// Right Up

	// Left down
	// Right down
	sum := 0
	if aI == 0 || aJ == 0 || aI == len(file)-1 || aJ == len(file[aI])-1 {
		return 0
	}
	if file[aI-1][aJ-1] == rune('M') && file[aI+1][aJ+1] == rune('S') || file[aI+1][aJ+1] == rune('M') && file[aI-1][aJ-1] == rune('S') {
		if file[aI-1][aJ+1] == rune('M') && file[aI+1][aJ-1] == rune('S') || file[aI+1][aJ-1] == rune('M') && file[aI-1][aJ+1] == rune('S') {
			sum++
		}
	}

	return sum
}

func ceresSearch(f *os.File) (string, error) {
	file := getFile(f)

	sum := 0
	for i, line := range file {
		for j, char := range line {
			if char == rune('X') {
				for range findXmas(file, i, j) {
					sum++
				}
			}
		}
	}
	return fmt.Sprint(sum), nil
}

func findXmas(file [][]rune, xI int, xJ int) [][][]int {
	var res [][][]int
	left := verifyDirection(file, xI, xJ, 0, -1)
	if left != nil {
		res = append(res, left)
	}
	right := verifyDirection(file, xI, xJ, 0, 1)
	if right != nil {
		res = append(res, right)
	}
	up := verifyDirection(file, xI, xJ, -1, 0)
	if up != nil {
		res = append(res, up)
	}
	down := verifyDirection(file, xI, xJ, 1, 0)
	if down != nil {
		res = append(res, down)
	}
	upLeft := verifyDirection(file, xI, xJ, -1, -1)
	if upLeft != nil {
		res = append(res, upLeft)
	}
	upRight := verifyDirection(file, xI, xJ, -1, 1)
	if upRight != nil {
		res = append(res, upRight)
	}
	downLeft := verifyDirection(file, xI, xJ, 1, -1)
	if downLeft != nil {
		res = append(res, downLeft)
	}
	downRight := verifyDirection(file, xI, xJ, 1, 1)
	if downRight != nil {
		res = append(res, downRight)
	}
	return res
}

func verifyDirection(file [][]rune, xI, xJ, dI, dJ int) [][]int {
	expectedString := "XMAS"
	var res [][]int
	for i, expected := range expectedString {
		curI := xI + (i * dI)
		curJ := xJ + (i * dJ)
		if curI < 0 || curJ < 0 || curI >= len(file) || curJ >= len(file[curI]) || expected != file[curI][curJ] {
			return nil
		}
		res = append(res, []int{curI, curJ})
	}
	return res
}

func getFile(f *os.File) [][]rune {
	scanner := bufio.NewScanner(f)
	var res [][]rune
	for scanner.Scan() {
		var l []rune
		for _, r := range scanner.Text() {
			l = append(l, r)
		}
		res = append(res, l)
	}

	return res
}
