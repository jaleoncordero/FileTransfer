package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/jaleoncordero/worker"
)

var (
	rgx    *regexp.Regexp
	dstDir string
)

func Run() {
	fmt.Println()

	err := validateArguments()
	if err != nil {
		panic(err.Error())
	}

	// set destination directory global
	dstDir, err = filepath.Abs(filepath.Clean(os.Args[2]))
	if err != nil {
		panic(err.Error())
	}

	err = os.MkdirAll(dstDir, 0777)
	if err != nil {
		panic(err.Error())
	}

	rgx, err = regexp.Compile(imgRegex)
	if err != nil {
		panic("failed to compile image regex")
	}

	wp := worker.NewPool(workerPoolSize)
	wp.Start()

	err = iterateDirectories(os.Args[1], &wp)
	if err != nil {
		panic(err.Error())
	}

	err = wp.Close()
	if err != nil {
		panic(err.Error())
	}
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
			iterateDirectories(filepath.Join(currentDir, file.Name()), wp)
		}
	}

	return nil
}
