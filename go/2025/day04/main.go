package main

import (
	"bufio"
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
	err = app.Run(printingDepartment)
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}
}


func printingDepartment(f *os.File) (string, error) {
    m := parseMap(f)

    sum := 0
    for rowNum, row := range m {
        for colNum, char := range row {
            if char == '@' && isAccessible(m, rowNum, colNum) {
                sum ++
            }
        }
    }
    
    return strconv.Itoa(sum), nil
}

func isAccessible(m [][]rune, rowNum, colNum int) bool {
    count := 0
    for i := rowNum - 1; i <= rowNum +1; i ++ {
        if i < 0 || i >= len(m) {
            continue
        }
        for j := colNum - 1; j <= colNum +1; j++ {
            if j < 0 || j >= len(m[i]) || (i == rowNum && j == colNum) {
                continue
            }

            if m[i][j] == '@'{
                count ++
            }
        }
    }
    return count < 4
}

func parseMap(f *os.File) [][]rune{
    s := bufio.NewScanner(f)
    
    m := [][]rune{}

    row := 0
    for s.Scan() {
        m = append(m, []rune{})
        for _, char := range s.Text() {
            m[row] = append(m[row], char)
        }
        row ++
    }
    return m
}
