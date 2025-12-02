package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"slices"
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
	err = app.Run(redNosedReport)
	if err != nil {
		logger.Error(err.Error())
	}

	err = app.Run(redNosedReport2)
	if err != nil {
		logger.Error(err.Error())
	}
}

func redNosedReport(f *os.File) (string, error) {
	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		report, err := parseLine(scanner.Text())
		if err != nil {
			return "", err
		}
		if isValidReport(report) {
			counter++
		}
	}
	return fmt.Sprint(counter), nil
}

func isValidReport(report []int) bool {
	// The levels are either all increasing or all decreasing.
	// Any two adjacent levels differ by at least one and at most three.
	direction := report[0] - report[1]
	for i := 0; i < len(report)-1; i++ {
		if !validatePair(direction, report[i], report[i+1]) {
			return false
		}
	}
	return true
}

func redNosedReport2(f *os.File) (string, error) {
	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		report, err := parseLine(scanner.Text())
		if err != nil {
			return "", err
		}
		if isValidReportWithRemoval(report) {
			counter++
		}
	}
	return fmt.Sprint(counter), nil
}

func isValidReportWithRemoval(report []int) bool {
	if isValidReport(report) {
		return true
	}
	for i := range report {
		cl := slices.Clone(report)
		rm := slices.Delete(cl, i, i+1)
		if isValidReport(rm) {
			return true
		}
	}
    return false
}

func validatePair(direction, i, j int) bool {
	if i-j == 0 {
		return false
	}

	if i-j > 3 || i-j < -3 {
		return false
	}
	if direction < 0 && i-j >= 0 {
		return false
	}
	if direction > 0 && i-j <= 0 {
		return false
	}
	return true
}

func parseLine(line string) ([]int, error) {
	split := strings.Split(line, " ")
	res := make([]int, len(split))
	for i, s := range split {
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		res[i] = val
	}
	return res, nil
}
