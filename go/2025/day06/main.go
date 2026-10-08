package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/YourAverageMoron/aoc/lib/app"
)

func main() {
	logger := slog.New(slog.Default().Handler())
	app, err := app.New(logger)
	if err != nil {
		logger.Error("error initialising app", slog.Any("error", err))
		return
	}
	err = app.Run(getTrashCompactor(getCalcListsPart1))
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}
	err = app.Run(getTrashCompactor(getCalcListPart2))
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}
}

func getTrashCompactor(parseFile func(*os.File) [][]string) func(f *os.File) (string, error) {
	return func(f *os.File) (string, error) {
		cl := parseFile(f)
		sum := 0
		for _, c := range cl {
			val, err := strconv.Atoi(c[0])
			if err != nil {
				return "", fmt.Errorf("failed to parse string to int (%s): %w", c[0], err)
			}
			for i := 1; i < len(c)-1; i++ {
				curr, err := strconv.Atoi(c[i])
				if err != nil {
					return "", fmt.Errorf("failed to parse string to int (%s): %w", c[i], err)
				}
				if c[len(c)-1] == "*" {
					val *= curr
				}
				if c[len(c)-1] == "+" {
					val += curr
				}
			}
			sum += val
		}
		return strconv.Itoa(sum), nil
	}
}

func getCalcListsPart1(f *os.File) [][]string {
	sc := bufio.NewScanner(f)

	res := [][]string{}

	for sc.Scan() {
		sb := strings.Builder{}
		col := 0
		for _, char := range sc.Text() {
			if char == ' ' {
				res, col = writeString(sb.String(), col, res)
				sb.Reset()
				continue
			}
			sb.WriteRune(char)
		}
		res, col = writeString(sb.String(), col, res)
	}
	return res
}

func writeString(s string, col int, listOLists [][]string) ([][]string, int) {
	if len(s) != 0 {
		if len(listOLists) <= col {
			listOLists = append(listOLists, []string{})
		}
		listOLists[col] = append(listOLists[col], s)
		col++
	}
	return listOLists, col
}


func getCalcListPart2(f *os.File) [][]string {
	fl := [][]rune{}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		row := []rune{}
		for _, r := range sc.Text() {
			row = append(row, r)
		}
		fl = append(fl, row)
	}

	calcs := [][]string{}
	curr := []string{}
	for j := len(fl[0]) - 1; j >= 0; j-- {
		sb := strings.Builder{}
		for i := 0; i < len(fl)-1; i++ {
			if fl[i][j] != ' ' {
				sb.WriteRune(fl[i][j])
			}
		}
		if sb.Len() > 0 {
			curr = append(curr, sb.String())
			sb.Reset()
		}

		if fl[len(fl)-1][j] != ' ' {
			curr = append(curr, string(fl[len(fl)-1][j]))
			calcs = append(calcs, curr)
			curr = []string{}
		}
	}

	return calcs
}
