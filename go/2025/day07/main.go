package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/YourAverageMoron/aoc/lib/app"
)

func main() {
	logger := slog.New(slog.Default().Handler())
	app, err := app.New(logger)
	if err != nil {
		logger.Error("error initialising app", slog.Any("error", err))
		return
	}
	err = app.Run(laboratories)
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}
}

func laboratories(f *os.File) (string, error) {
	sc := bufio.NewScanner(f)
	sc.Scan()
	prev := sc.Bytes()
	count := 0
	for sc.Scan() {
		curr := sc.Bytes()
		for i, b := range prev {
			if b == '|' || b == 'S' {
				if curr[i] == '^' {
					count++
					if i > 0 {
						curr[i-1] = '|'
					}
					if i < len(curr)-1 {
						curr[i+1] = '|'
					}
				} else {
					curr[i] = '|'
				}
			}
		}
        prev = make([]byte, len(prev))
		copy(prev, curr)
	}
	return strconv.Itoa(count), nil
}

func printRow(r []byte){
    for _, b := range r{
        fmt.Print(string(b))
    }
    fmt.Println()
}
