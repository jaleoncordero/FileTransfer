package internal

import (
	"fmt"
	"os"
	"regexp"
)

var (
	rgx *regexp.Regexp

	srcDir, dstDir string
)

func Init() {
	r, err := regexp.Compile(imgRegex)

	if err != nil {
		panic("failed to compile image regex")
	}
	rgx = r
}

func Run() {
	if err := validateArguments(); err != nil {
		panic(err.Error())
	}

}

func validateArguments() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("missing 1 or more arguments")
	}

	// source directory is first argument
	if err := validateDirectoryExists("source", os.Args[1]); err != nil {
		return err
	} else {
		srcDir = os.Args[1]
	}

	// destination directory is second argument
	if err := validateDirectoryExists("source", os.Args[2]); err != nil {
		return err
	} else {
		dstDir = os.Args[2]
	}

	return nil
}

func validateDirectoryExists(dirType, d string) error {
	fileInfo, err := os.Stat(d)
	if err != nil {
		return err
	}

	fmt.Printf("%s directory information\n============================\n%v\n\n", dirType, fileInfo)
	return nil
}
