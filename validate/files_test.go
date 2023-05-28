package validate

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func Test_FileExists(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		setupFile  func(t *testing.T) (filePath string)
		errWrapped error
		errRegex   *regexp.Regexp
	}{
		"empty_path": {
			setupFile: func(t *testing.T) (filePath string) {
				t.Helper()
				return ""
			},
			errWrapped: ErrFilePathIsDir,
			errRegex:   regexp.MustCompile(`^filepath is a directory: \.$`),
		},
		"directory": {
			setupFile: func(t *testing.T) (filePath string) {
				t.Helper()
				path := filepath.Join(t.TempDir(), "directory")
				err := os.MkdirAll(path, os.ModePerm)
				if err != nil {
					t.Fatal(err)
				}
				return path
			},
			errWrapped: ErrFilePathIsDir,
			errRegex:   regexp.MustCompile(`^filepath is a directory: /[a-zA-Z0-9_./]+/directory$`),
		},
		"file_not_exists": {
			setupFile: func(t *testing.T) (filePath string) {
				t.Helper()
				return filepath.Join(t.TempDir(), "file.txt")
			},
			errWrapped: ErrFileDoesNotExist,
			errRegex:   regexp.MustCompile("file does not exist: /[a-zA-Z0-9_./]+/file.txt"),
		},
		"file_exists": {
			setupFile: func(t *testing.T) (filePath string) {
				t.Helper()
				path := filepath.Join(t.TempDir(), "file.txt")
				_, err := os.Create(path)
				if err != nil {
					t.Fatal(err)
				}
				return path
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			filePath := testCase.setupFile(t)

			err := FileExists(filePath)

			if !errors.Is(err, testCase.errWrapped) {
				t.Errorf("expected error '%v' to be wrapped, but it is not in '%v'", testCase.errWrapped, err)
			}
			if testCase.errWrapped != nil &&
				!testCase.errRegex.MatchString(fmt.Sprint(err)) {
				t.Errorf("expected error message '%v' to match regex %s",
					err, testCase.errRegex)
			}
		})
	}
}
