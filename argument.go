package arguments

import (
	"os"
)

type Options[T any] struct {
	Min T
	Max T
}

type argument struct {
	isNil bool
	value string
}

func asFile(path string, flag int, perm os.FileMode, def *os.File) (*os.File, error) {
	if path == "" {
		return def, nil
	}

	f, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}
