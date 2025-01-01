package internal

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
)

type CopyFileJob struct {
	srcDir string
}

// implement
func (j *CopyFileJob) Process() error {
	files, err := os.ReadDir(j.srcDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filename := file.Name()

		// if item is an image file, we want to process it
		if !file.IsDir() && rgx.MatchString(filename) {
			sp := filepath.Join(j.srcDir, filename)
			dp := filepath.Join(dstDir, filename)

			if err := copyFile(sp, dp); err != nil {
				return err
			}

			fileMu.Lock()
			totalFiles += 1
			fileMu.Unlock()
		}
	}

	return nil
}

func copyFile(srcPath, outPath string) error {
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

	out, err := os.Create(outPath)
	if err != nil {
		return err
	}

	defer func() {
		if e := out.Close(); e != nil {
			err = errors.Join(err, e)
		}
	}()

	_, err = io.Copy(io.MultiWriter(out, bar), reader)
	if err != nil {
		return err
	}

	return nil
}
