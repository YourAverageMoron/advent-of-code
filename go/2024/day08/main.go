package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"math"
	"os"

	"github.com/YourAverageMoron/aoc/lib/app"
)

func main() {
	logger := slog.New(slog.Default().Handler())
	app, err := app.New(logger)
	if err != nil {
		logger.Error("error initialising app", slog.Any("error", err))
		return
	}
	err = app.Run(resonantCollinearity)
	if err != nil {
		logger.Error(err.Error())
	}
	// err = app.Run(temp)
	// if err != nil {
	// 	logger.Error(err.Error())
	// }
}

type antennaMap struct {
	m    map[rune][][]int
	iMax int
	jMax int
}

func resonantCollinearity(f *os.File) (string, error) {
	m := parseMap(f)


    sum := 0
	a := getAntinodes(m)

    for  range a {
        sum ++
    }
	return fmt.Sprint(sum), nil
}

func getAntinodes(m *antennaMap) map[string]bool {
    res := map[string]bool{}
	for _, l := range m.m {
		for _, c1 := range l {
			for _, c2 := range l {
				as := calculateAntennas(c1, c2)
				for _, a := range as {
					if isInMap(m.iMax, m.jMax, a) {
                        res[fmt.Sprintf("%d:%d", a[0], a[1])] = true
					}
				}
			}
		}
	}
	return res
}

func isInMap(iMax, jMax int, a []int) bool {
	return a[0] >= 0 && a[0] <= iMax && a[1] >= 0 && a[1] <= jMax
}

func calculateAntennas(c1 []int, c2 []int) [][]int {
	iDiff := int(math.Abs(float64(c1[0] - c2[0])))
	jDiff := int(math.Abs(float64(c1[1] - c2[1])))

	a1 := make([]int, 2)
	a2 := make([]int, 2)
	if c1[0] > c2[0] {
		a1[0] = c1[0] + iDiff
		a2[0] = c2[0] - iDiff
	} else {
		a1[0] = c1[0] - iDiff
		a2[0] = c2[0] + iDiff
	}
	if c1[1] > c2[1] {
		a1[1] = c1[1] + jDiff
		a2[1] = c2[1] - jDiff
	} else {
		a1[1] = c1[1] - jDiff
		a2[1] = c2[1] + jDiff
	}
	return [][]int{a1, a2}
}

func parseMap(f *os.File) *antennaMap {
	m := map[rune][][]int{}
	buff := bufio.NewScanner(f)
	var iMax int
	var jMax int
	for i := 0; buff.Scan(); i++ {
		iMax = i
		l := buff.Text()
		for j, r := range l {
			if r == rune(46) {
				continue
			}
			m[r] = append(m[r], []int{i, j})
			jMax = j
		}
	}
	return &antennaMap{
		iMax: iMax,
		jMax: jMax,
		m:    m,
	}
}
