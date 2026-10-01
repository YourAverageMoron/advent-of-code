package main

import (
	"bufio"
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
	err = app.Run(getLobby(getMaxJoltagePart1))
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}

	err = app.Run(getLobby(getMaxJoltagePart2))
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}

}

func getLobby(getMaxJoltage func(string) (int, error)) func(f *os.File) (string, error) {
	return func(f *os.File) (string, error) {
		sc := bufio.NewScanner(f)
		sum := 0
		for sc.Scan() {
			j, err := getMaxJoltage(sc.Text())
			if err != nil {
				return "", err
			}
			sum += j
		}

		return strconv.Itoa(sum), nil
	}
}

func getMaxJoltagePart1(bank string) (int, error) {
	l := 0
	r := 0
	for i, c := range bank {
		if int(c-'0') > r {
			r = int(c - '0')
		}
		if int(c-'0') > l && i < len(bank)-1 {
			l = int(c - '0')
			r = 0
		}
	}

	strRes := strconv.Itoa(l) + strconv.Itoa(r)
	return strconv.Atoi(strRes)
}

func getMaxJoltagePart2(bank string) (int, error) {
	arr := make([]int, 12)
	for i, char := range bank {
		startPos := 12 - len(bank) + i
        if startPos < 0 {
                startPos = 0
            }
        for curPos := startPos; curPos < 12; curPos ++ {
            if int(char-'0') > arr[curPos] {
                arr[curPos] = int(char - '0')
                for i := curPos + 1; i < len(arr); i++ {
                    arr[i] = 0
                }
                break
            }
        }
	}

    sb := strings.Builder{}
    for _, v := range arr {
        sb.Write([]byte(strconv.Itoa(v)))
    }
    return strconv.Atoi(sb.String())
}
