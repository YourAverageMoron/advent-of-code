package app

import (
	"fmt"
	"os"
	"path"

	"github.com/YourAverageMoron/aoc/go/lib/args"
)

type logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

type callbackFunc func(*os.File) (int, error)

type App struct {
	logger logger
	file   string
}

func New(logger logger) (*App, error) {
	args, err := args.Parse(logger)
	if err != nil {
		return nil, err
	}
	return &App{
		file:   getDataFilePath(logger, args.DataPath),
		logger: logger,
	}, nil
}

func (a *App) Run(callbackFunc callbackFunc) error {
	f, err := os.Open(a.file)
	if err != nil {
		return err
	}
	defer f.Close()
	res, err := callbackFunc(f)
	if err != nil {
		return err
	}
	a.logger.Info(fmt.Sprintf("result - %d", res))
	return nil
}

func getDataFilePath(logger logger, dataPath string) string {
	demoEnv := os.Getenv("DEMO")
	var file string
	if demoEnv != "" {
		logger.Info("running demo")
		file = "input_demo.txt"
	} else {
		logger.Info("running prod")
		file = "input.txt"
	}
	return path.Join(dataPath, file)
}
