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
	err = app.Run(trashCompactor)
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}
}

func trashCompactor(f *os.File) (string, error) {

	cl := getCalcLists(f)

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

func getCalcLists(f *os.File) [][]string {
	sc := bufio.NewScanner(f)
	sb := strings.Builder{}

	res := [][]string{}
	col := 0

	for sc.Scan() {
		for _, char := range sc.Text() {
			if char == ' ' {
				res, col, sb = writeString(sb, col, res)

				continue
			}
			sb.WriteRune(char)
		}
		res, col, sb = writeString(sb, col, res)
		col = 0
	}
	return res
}

func writeString(sb strings.Builder, col int, listOLists [][]string) ([][]string, int, strings.Builder) {
	if sb.Len() != 0 {
		if len(listOLists) <= col {
			listOLists = append(listOLists, []string{})
		}
		s := sb.String()
		listOLists[col] = append(listOLists[col], s)
		sb.Reset()
		col++
	}
	return listOLists, col, sb
}
