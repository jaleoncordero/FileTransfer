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

func buildExtensionRegex() (string, error) {
	r := matchRegex
	first := true

	l, err := buildExtensionList()
	if err != nil {
		return "", err
	}

	for _, ext := range l {
		if first {
			first = false
		} else {
			r += "|"
		}

		r += fmt.Sprintf("(.*%s$)", ext)
	}

	return r, nil
}

func buildExtensionList() ([]string, error) {
	fmt.Printf("MODE: %s", execMode)
	if v, ok := supportedExtensions[execMode]; ok {
		return v, nil
	}

	// TODO: support custom mode
	switch execMode {
	default:
		return nil, fmt.Errorf("mode %s not supported", execMode)
	}
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
