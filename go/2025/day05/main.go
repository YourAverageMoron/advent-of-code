package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"sort"
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
	err = app.Run(getCafeteria(binarySearchRanges))
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}
	err = app.Run(getCafeteria(countAvailableIngredients))
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}
}

func getCafeteria(calc func(rl []*freshRange, ids []string) (string, error)) func(f *os.File) (string, error) {
	return func(f *os.File) (string, error) {
		rangeStrings, ids := parseFile(f)
		rl, err := createFreshRangeList(rangeStrings)
		if err != nil {
			return "", fmt.Errorf("failed to create fresh map: %w", err)
		}
		return calc(rl, ids)
	}
}

func binarySearchRanges(rl []*freshRange, ids []string) (string, error) {
	count := 0
	for _, id := range ids {
		iid, err := strconv.Atoi(id)
		if err != nil {
			return "", fmt.Errorf("failed to parse if (%s): %w", err)
		}
		if binarySearch(iid, rl) {
			count++
		}
	}

	return strconv.Itoa(count), nil
}

// 350684792662863 - TO HIGH
func countAvailableIngredients(rl []*freshRange, ids []string) (string, error) {
	sum := 0
	for _, r := range rl {
		sum += r.end - r.start + 1
	}
	return strconv.Itoa(sum), nil
}

func binarySearch(value int, rl []*freshRange) bool {
	l := 0
	r := len(rl) - 1

	for l <= r {
		mid := (r + l) / 2
		if rl[mid].start <= value && rl[mid].end >= value {
			return true
		}
		if value > rl[mid].end {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

type freshRange struct {
	start int
	end   int
}

func createFreshRangeList(rangeStrings []string) ([]*freshRange, error) {
	rl := make([]*freshRange, 0, len(rangeStrings))
	for _, rs := range rangeStrings {
		splitRs := strings.Split(rs, "-")
		if len(splitRs) != 2 {
			return nil, fmt.Errorf("invalid range string (%s)", rs)
		}
		start, err := strconv.Atoi(splitRs[0])
		if err != nil {
			return nil, fmt.Errorf("unable to parse start range (%s): %w", splitRs[0], err)
		}
		end, err := strconv.Atoi(splitRs[1])
		if err != nil {
			return nil, fmt.Errorf("unable to parse end range (%s): %w", splitRs[1], err)
		}
		rl = append(rl, &freshRange{start: start, end: end})
	}

	sort.Slice(rl, func(i, j int) bool {
		return rl[i].start < rl[j].start
	})

	for i, r := range rl {
		for i+1 < len(rl) && r.end >= rl[i+1].start {
			if r.end < rl[i+1].end {
				r.end = rl[i+1].end
			}
			rl = append(rl[:i+1], rl[i+2:]...)
		}
	}
	return rl, nil
}

func parseFile(f *os.File) ([]string, []string) {
	sc := bufio.NewScanner(f)
	ranges := []string{}
	ids := []string{}

	hitSpace := false
	for sc.Scan() {
		t := sc.Text()
		if t == "" {
			hitSpace = true
			continue
		}
		if !hitSpace {
			ranges = append(ranges, t)
		} else {
			ids = append(ids, t)
		}
	}

	return ranges, ids
}
