package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
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
	err = app.Run(printQueue)
	if err != nil {
		logger.Error(err.Error())
	}

}

func printQueue(f *os.File) (string, error) {
	rules, updates := parseFile(f)

	rulesMap := createRulesMap(rules)

	sum := 0
	for _, u := range updates {
		s := strings.Split(u, ",")
		if checkUpdate(s, rulesMap) {
			v, err := strconv.Atoi(s[len(s)/2])
			if err != nil {
                return "", err
			}
			sum += v
		}
	}

	return fmt.Sprint(sum), nil
}

func checkUpdate(s []string, m map[string]map[string]bool) bool {
	for i, curr := range s {
		for _, check := range s[i+1:] {
			if !m[curr][check] {
				return false
			}
		}
	}
	return true
}

func createRulesMap(rules []string) map[string]map[string]bool {
	res := make(map[string]map[string]bool)
	for _, rule := range rules {
		split := strings.Split(rule, "|")
		_, ok := res[split[0]]
		if !ok {
			res[split[0]] = map[string]bool{}
		}
		res[split[0]][split[1]] = true
	}
	return res
}

func parseFile(r io.Reader) ([]string, []string) {
	scanner := bufio.NewScanner(r)
	hitUpdates := false
	var rules, updates []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			hitUpdates = true
			continue
		}
		if hitUpdates {
			updates = append(updates, scanner.Text())
		} else {
			rules = append(rules, scanner.Text())
		}
	}

	return rules, updates
}
