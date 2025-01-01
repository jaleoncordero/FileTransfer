package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/jaleoncordero/worker"
	"github.com/schollz/progressbar/v3"
)

func Run() {
	fmt.Println()

	// validate command input
	err := validateArguments()
	if err != nil {
		panic(err.Error())
	}

	// compile regex that determines what media files to copy
	rgx, err = regexp.Compile(getRegex())
	if err != nil {
		panic("failed to compile media regex")
	}

	// start worker pool
	wp := worker.NewPool(workerPoolSize)
	wp.Start()

	// progress bar :)
	bar = progressbar.DefaultBytes(
		-1,
		"copying files",
	)

	// begin work
	err = iterateDirectories(srcDir, &wp)
	if err != nil {
		panic(err.Error())
	}

	// wrap up work
	err = wp.Close()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("\n\nSuccessfully copied %d file(s) to %s\n", totalFiles, dstDir)
}

func iterateDirectories(currentDir string, wp *worker.Pool) error {
	files, err := os.ReadDir(currentDir)
	if err != nil {
		return err
	}

	// add file copy job to worker pool
	wp.AddJob(
		&CopyFileJob{
			srcDir: currentDir,
		},
	)

	// iterate through items in current directory
	for _, file := range files {

		// if item is a directory, we want to iterate into it
		if file.IsDir() {
			err := iterateDirectories(filepath.Join(currentDir, file.Name()), wp)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
