package internal

import (
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
		fp := filepath.Join(j.srcDir, filename)

		// if item is an image file, we want to process it
		if !file.IsDir() && rgx.MatchString(filename) {

			data, err := os.ReadFile(fp)
			if err != nil {
				return err
			}

			err = os.WriteFile(filepath.Join(dstDir, filename), data, 0777) //write the content to destination file
			if err != nil {
				return err
			}
		}
	}

	return nil
}
