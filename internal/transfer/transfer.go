package transfer

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
	srcDir, dstDir, err := validateArguments()
	if err != nil {
		panic(err.Error())
	}

	if execMode != "all" {

		// build & compile regex that determines what media files to copy
		rS, err := buildExtensionRegex()
		if err != nil {
			panic(err.Error())
		}

		rgx, err = regexp.Compile(rS)
		if err != nil {
			panic(err.Error())
		}
	}

	// start worker pool
	wp := worker.NewPool(workerPoolSize)
	wp.Start()

	// progress bar :)
	pb = progressbar.DefaultBytes(
		-1,
		"copying files",
	)

	// begin work
	err = iterateDirectories(srcDir, dstDir, &wp)
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

func iterateDirectories(currentDir, dstDir string, wp *worker.Pool) error {
	files, err := os.ReadDir(currentDir)
	if err != nil {
		return err
	}

	// add file copy job to worker pool - each worker copies files from current dir
	wp.AddJob(
		&CopyFileJob{
			dstDir:   dstDir,
			srcDir:   currentDir,
			useRegex: rgx != nil,
			rgx:      rgx,
		},
	)

	// iterate through items in current directory
	for _, file := range files {

		// if item is a directory, we want to iterate into it
		if file.IsDir() {
			err := iterateDirectories(filepath.Join(currentDir, file.Name()), dstDir, wp)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
