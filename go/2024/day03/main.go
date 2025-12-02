package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
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
	err = app.Run(mulItOver)
	if err != nil {
		logger.Error(err.Error())
	}
	err = app.Run(mulItOver2)
	if err != nil {
		logger.Error(err.Error())
	}
}

func mulItOver(f *os.File) (string, error) {
	scanner := bufio.NewScanner(f)
	s := ""
	for scanner.Scan() {
		s += scanner.Text()
	}
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := re.FindAll([]byte(s), -1)

	sum := 0
	for _, match := range matches {
        val, err := calculateMul(match)
        if err != nil{
            return "", err
        }
        sum += val
	}

	return fmt.Sprint(sum), nil
}

func mulItOver2(f *os.File) (string, error) {
	scanner := bufio.NewScanner(f)
	s := ""
	for scanner.Scan() {
		s += scanner.Text()
	}
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)

    sum := 0
    do := true
	matches := re.FindAll([]byte(s), -1)
	for _, match := range matches {
        if string(match) == "do()" {
            do = true
        } else if string(match) == "don't()" {
            do = false
        } else {
            if !do {
                continue
            }
            val, err := calculateMul(match)
            if err != nil {
                return "", err
            }
            sum += val
        }
	}
	return fmt.Sprint(sum), nil
}

func calculateMul(match []byte) (int, error){
	stringNums := strings.Split(string(match[4:len(match)-1]), ",")
	nums := make([]int, len(stringNums))
	for i, val := range stringNums {
		num, err := strconv.Atoi(val)
		if err != nil {
			return 0, err
		}
		nums[i] = num
	}
    return nums[0] * nums[1], nil
}
