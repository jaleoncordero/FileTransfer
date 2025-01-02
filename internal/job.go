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

		// if item is an matching file type, we want to process it
		if !file.IsDir() && rgx.MatchString(filename) {
			sp := filepath.Join(j.srcDir, filename)

			// handle duplicate destination filenames. No duplicate file
			// check is made
			dp, err := getUniqueFilename(dstDir, filename, 1)
			if err != nil {
				return err
			}

			if err := copyFile(sp, dp); err != nil {
				return err
			}

			// increase file counter
			fileMu.Lock()
			totalFiles += 1
			fileMu.Unlock()
		}
	}

	return nil
}
