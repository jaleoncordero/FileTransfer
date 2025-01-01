package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/inhies/go-bytesize"
)

func validateArguments() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("missing 1 or more arguments")
	}

	// source directory is first argument
	if err := validateDirectoryExists("source", os.Args[1]); err != nil {
		return err
	}

	return nil
}

func validateDirectoryExists(dirType, d string) error {
	path, err := filepath.Abs(filepath.Clean(d))
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	fmt.Printf("%s directory information\n----------------------------\nPath: %s\nSize: %v\n\n",
		dirType, d, bytesize.New(float64(fileInfo.Size())))

	return nil
}
