package main

import (
	"bufio"
	"fmt"
	"io"
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
	err = app.Run(GetGiftShopFunc(isValidPart1))
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}

    err = app.Run(GetGiftShopFunc(isValidPart2))
	if err != nil {
		logger.Error("error running app", slog.Any("error", err))
		return
	}

}


func GetGiftShopFunc(isValid func (int) bool) func (f *os.File) (string, error) {
    return func (f *os.File) (string, error) {
	ch := make(chan string)
	go readIdRanges(f, ch)

	sum := 0

	for idRange := range ch {
		start, end, err := parseRange(idRange)
		if err != nil {
			return "", err
		}
		for i := start; i <= end; i++ {
			if !isValid(i) {
				sum += i
			}
		}
	}

	return strconv.Itoa(sum), nil
}
}

func isValidPart1(input int) bool {
	si := strconv.Itoa(input)
	if len(si)%2 != 0 {
		return true
	}

	for i := 0; i < len(si)/2; i++ {
		if si[i] != si[len(si)/2+i] {
			return true
		}
	}
	return false
}

func isValidPart2(input int) bool {
	si := strconv.Itoa(input)
    for subLen := 1; subLen <= len(si) / 2; subLen ++ {
        if len(si) % subLen != 0 {
            // not divisible e.g "1231231" and sublen = 3 
            continue
        }
        if isRepeated(si, si[0:subLen]){
            return false
        }
    }
    return true
}

func isRepeated(input, substring string) bool {
    for i:= 0; i + len(substring) <= len(input); i+= len(substring) {
        if input[i:i+len(substring)] != substring{
            return false
        } 
    }
    return true
}

func parseRange(rId string) (int, int, error) {
	splitRId := strings.Split(rId, "-")
	if len(splitRId) != 2 {
		return 0, 0, fmt.Errorf("invalid rId: %s", rId)
	}
	start, err := strconv.Atoi(splitRId[0])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse start id (%s): %w", splitRId[0], err)
	}
	end, err := strconv.Atoi(splitRId[1])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse end id (%s): %w", splitRId[1], err)
	}
	return start, end, nil
}

func readIdRanges(f *os.File, out chan<- string) error {
	r := bufio.NewReader(f)

	sb := strings.Builder{}
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			// TODO: log this out better i think?
			if err == io.EOF {
				out <- sb.String()
				sb.Reset()
				break
			}
			fmt.Println("error:", err)
			return err
		}

		if c == ',' {
			out <- sb.String()
			sb.Reset()
			continue
		}
		if c == '-' || (c >= '0' && c <= '9') {
			sb.WriteRune(c)
		}
	}
	close(out)
	return nil
}
