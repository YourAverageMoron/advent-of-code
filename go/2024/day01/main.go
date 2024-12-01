package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/YourAverageMoron/aoc/go/lib/app"
)

func main() {
	logger := slog.New(slog.Default().Handler())
	app, err := app.New(logger)
	if err != nil {
		logger.Error("error initialising app", slog.Any("error", err))
		return
	}
	err = app.Run(historianHysteria)
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}

	err = app.Run(historianHysteria2)
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}
}

func parseLine(line string) (int, int, error) {
	split := strings.Split(line, "   ")
	if len(split) != 2 {
		return -1, -1, fmt.Errorf("unable to parse line - incorrect split seperator\n")
	}
	i1, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, err
	}
	i2, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, err
	}
	return i1, i2, nil
}

func historianHysteria(f *os.File) (int, error) {
	scanner := bufio.NewScanner(f)
	var s1, s2 []int
	for scanner.Scan() {
		i1, i2, err := parseLine(scanner.Text())
		if err != nil {
			return 0, err
		}
		s1 = append(s1, i1)
		s2 = append(s2, i2)
	}
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})

	sum := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] > s2[i] {
			sum += s1[i] - s2[i]
		} else {
			sum += s2[i] - s1[i]
		}
	}
	return sum, nil
}

func historianHysteria2(f *os.File) (int, error) {
	scanner := bufio.NewScanner(f)
	var s []int
	m := map[int]int{}
	for scanner.Scan() {
		i1, i2, err := parseLine(scanner.Text())
		if err != nil {
			return 0, err
		}
		s = append(s, i1)
		m[i2]++
	}
	sum := 0
	for _, i := range s {
		count, ok := m[i]
		if !ok {
			count = 0
		}
		sum += (i * count)
	}
	return sum, nil
}
