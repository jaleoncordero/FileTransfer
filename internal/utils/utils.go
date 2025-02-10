package utils

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func GetUniqueFilename(path, filename string, copy int) (string, error) {
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

func CopyFile(srcPath, outPath string, pb *progressbar.ProgressBar) error {
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

	// write source file to destination file, & to progress bar if any
	if pb != nil {
		_, err = io.Copy(io.MultiWriter(out, pb), reader)
		if err != nil {
			return err
		}
	} else {
		_, err = io.Copy(out, reader)
		if err != nil {
			return err
		}
	}

	return nil
}

func ValidateDirectoryExists(d string) error {
	_, err := os.Stat(d)
	if err != nil {
		return err
	}

	return nil
}
