package main

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"slices"
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
	err = app.Run(bridgeRepair)
	if err != nil {
		logger.Error(err.Error())
	}
	err = app.Run(bridgeRepair2)
	if err != nil {
		logger.Error(err.Error())
	}
}

func bridgeRepair2(f *os.File) (string, error) {
	buff := bufio.NewScanner(f)
	sum := 0
	for buff.Scan() {
		value, args, err := parseLine(buff.Text())
		if err != nil {
			return "", err
		}
		slices.Reverse(args)
		if isValidEquation2(value, args) {
			sum += value
		}
	}

	return fmt.Sprint(int(sum)), nil
}



func isValidEquation2(v int, args []int) bool {
	if len(args) == 1 {
		if v == args[0] {
			return true
		}
		return false
	}

    cat, err := strconv.Atoi( fmt.Sprintf("%d%d", args[len(args)-1], args[len(args)-2]))
    if err != nil {
        return false
    }

	if len(args) == 2 {
		if args[0]+args[1] == v || args[0]*args[1] == v || cat == v{
			return true
		}
		return false
	}

	newEndSum := args[len(args)-1] + args[len(args)-2]
	newEndProd := args[len(args)-1] * args[len(args)-2]

	argsSum := make([]int, len(args) - 1)
    copy(argsSum, args)
	argsSum[len(argsSum)-1] = newEndSum

	argsProd := make([]int, len(args) - 1)
    copy(argsProd, args)
	argsProd[len(argsProd)-1] = newEndProd

	argsCat := make([]int, len(args) - 1)
    copy(argsCat, args)
	argsCat[len(argsCat)-1] = cat


	if isValidEquation2(v, argsSum) || isValidEquation2(v, argsProd) || isValidEquation2(v, argsCat) {
		return true
	}
	return false
}



func bridgeRepair(f *os.File) (string, error) {
	buff := bufio.NewScanner(f)
	sum := 0
	for buff.Scan() {
		value, args, err := parseLine(buff.Text())
		if err != nil {
			return "", err
		}
		slices.Reverse(args)
		if isValidEquation(value, args) {
			sum += value
		}
	}

	return fmt.Sprint(int(sum)), nil
}



func isValidEquation(v int, args []int) bool {
	if len(args) == 1 {
		if v == args[0] {
			return true
		}
		return false
	}

	if len(args) == 2 {
		if args[0]+args[1] == v || args[0]*args[1] == v {
			return true
		}
		return false
	}

	newEndSum := args[len(args)-1] + args[len(args)-2]
	newEndProd := args[len(args)-1] * args[len(args)-2]

	argsSum := make([]int, len(args) - 1)
    copy(argsSum, args)
	argsSum[len(argsSum)-1] = newEndSum

	argsProd := make([]int, len(args) - 1)
    copy(argsProd, args)
	argsProd[len(argsProd)-1] = newEndProd

	if isValidEquation(v, argsSum) || isValidEquation(v, argsProd) {
		return true
	}
	return false
}

func parseLine(l string) (int, []int, error) {
	sp := strings.Split(l, ":")
	if len(sp) != 2 {
		return 0, nil, errors.New("not able to split by : - invalid length")
	}
	val, err := strconv.Atoi(sp[0])
	if err != nil {
		return 0, nil, err
	}
	argsS := strings.Split(sp[1], " ")
	args := make([]int, 0, len(argsS))
	for _, argS := range argsS {
		if argS == "" {
			continue
		}
		arg, err := strconv.Atoi(argS)
		if err != nil {
			return 0, nil, err
		}
		args = append(args, arg)
	}

	return val, args, nil
}
