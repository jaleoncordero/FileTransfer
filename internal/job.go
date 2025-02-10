package internal

import (
	"os"
	"path/filepath"
	"regexp"
)

type CopyFileJob struct {
	dstDir   string
	srcDir   string
	useRegex bool

	rgx *regexp.Regexp
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
		if !file.IsDir() && (!j.useRegex || j.rgx.MatchString(filename)) {
			sp := filepath.Join(j.srcDir, filename)

			// handle duplicate destination filenames. No duplicate file
			// check is made
			dp, err := getUniqueFilename(j.dstDir, filename, 1)
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
