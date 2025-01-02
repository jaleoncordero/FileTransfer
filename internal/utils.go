package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// TODO: pass in the different sets of extensions supported
func getRegex() string {
	r := matchRegex
	first := true

	for _, ext := range imageExtensions {
		if first {
			first = false
		} else {
			r += "|"
		}

		r += fmt.Sprintf("(.*%s$)", ext)
	}

	return r
}

func getUniqueFilename(path, filename string, copy int) (string, error) {
	for {
		dp := filepath.Join(path, filename)

		if _, err := os.Stat(dp); err == nil {
			// it exists, therefore we must append something to the filename to avoid overwriting
			s := strings.Split(filename, ".")
			filename = s[0] + "_collision_" + strconv.Itoa(copy) + "." + s[1]

			copy += 1
		} else if errors.Is(err, os.ErrNotExist) {
			// if it does not exist, then we can proceed
			return dp, nil

		} else {
			// some other error occurred
			return "", err
		}
	}
}

func copyFile(srcPath, outPath string) error {

	// open source file
	source, err := os.Open(srcPath)
	if err != nil {
		return err
	}

	defer func() {
		if e := source.Close(); e != nil {
			err = errors.Join(err, e)
		}
	}()

	reader := bufio.NewReader(source)

	// create destination file
	out, err := os.Create(outPath)
	if err != nil {
		return err
	}

	defer func() {
		if e := out.Close(); e != nil {
			err = errors.Join(err, e)
		}
	}()

	// write source file to destination file, & to progress bar
	_, err = io.Copy(io.MultiWriter(out, pb), reader)
	if err != nil {
		return err
	}

	return nil
}
