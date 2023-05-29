package validate

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrFileDoesNotExist = errors.New("file does not exist")
	ErrFileStat         = errors.New("failed stating file")
	ErrFilePathIsDir    = errors.New("filepath is a directory")
	ErrFilepathIsFile   = errors.New("filepath is a file")
)

func FileExists(path string) (err error) {
	const directory = false
	return fileExists(path, directory)
}

func DirectoryExists(path string) (err error) {
	const directory = true
	return fileExists(path, directory)
}

func fileExists(path string, directory bool) (err error) {
	path = filepath.Clean(path)

	stat, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("%w: %s", ErrFileDoesNotExist, path)
	} else if err != nil {
		return fmt.Errorf("%w: %w", ErrFileStat, err)
	}

	if directory && !stat.IsDir() {
		return fmt.Errorf("%w: %s", ErrFilepathIsFile, path)
	} else if !directory && stat.IsDir() {
		return fmt.Errorf("%w: %s", ErrFilePathIsDir, path)
	}

	return nil
}
