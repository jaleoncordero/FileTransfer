package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func validateArguments() (err error) {
	if len(os.Args) != 3 {
		return fmt.Errorf("missing 1 or more arguments")
	}

	// normalize source directory
	srcDir, err = filepath.Abs(filepath.Clean(os.Args[1]))
	if err != nil {
		return
	}

	if err = validateDirectoryExists(srcDir); err != nil {
		return
	}

	// normalize destination directory
	dstDir, err = filepath.Abs(filepath.Clean(os.Args[2]))
	if err != nil {
		return
	}

	// attempt to create destination directory if it doesn't exist
	if err = validateDirectoryExists(dstDir); os.IsNotExist(err) {
		if err = os.MkdirAll(dstDir, 0777); err != nil {
			return
		}

		return nil
	}

	return nil
}

func validateDirectoryExists(d string) error {
	_, err := os.Stat(d)
	if err != nil {
		return err
	}

	return nil
}
