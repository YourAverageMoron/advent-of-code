package args

import (
	"errors"
	"log/slog"
	"os"
)

type Args struct {
	DataPath string
}

type logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

func Parse(logger logger) (*Args, error) {
	args := os.Args
	if len(args) < 2 {
		err := errors.New("not enough arguments provided, expected 1, recieved 0")
		logger.Error("error parsing args", slog.Any("error", err))
		return nil, err
	}
	return &Args{
		DataPath: args[1],
	}, nil
}
