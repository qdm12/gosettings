package validate

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrFileDoesNotExist = errors.New("file does not exist")
	ErrFileOpen         = errors.New("failed opening file")
	ErrFileStat         = errors.New("failed stating file")
	ErrFilePathIsDir    = errors.New("filepath is a directory")
	ErrFileClose        = errors.New("failed closing file")
)

func FileExists(path string) (err error) {
	path = filepath.Clean(path)

	f, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("%w: %s", ErrFileDoesNotExist, path)
	} else if err != nil {
		return fmt.Errorf("%w: %w", ErrFileOpen, err)
	}

	stat, err := f.Stat()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFileStat, err)
	}

	if stat.IsDir() {
		return fmt.Errorf("%w: %s", ErrFilePathIsDir, path)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFileClose, err)
	}

	return nil
}
