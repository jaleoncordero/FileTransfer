package transfer

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/jaleoncordero/FileTransfer/internal/utils"
)

func validateArguments() (string, string, error) {
	if len(os.Args) < 3 {
		return "", "", fmt.Errorf("missing 1 or more arguments")
	}

	// normalize source directory
	srcDir, err := filepath.Abs(filepath.Clean(os.Args[1]))
	if err != nil {
		return "", "", err
	}

	// validate source dir
	if err = utils.ValidateDirectoryExists(srcDir); err != nil {
		return "", "", err
	}

	// normalize destination directory
	dstDir, err := filepath.Abs(filepath.Clean(os.Args[2]))
	if err != nil {
		return "", "", err
	}

	// attempt to create destination directory if it doesn't exist
	err = utils.ValidateDirectoryExists(dstDir)
	if errors.Is(err, fs.ErrNotExist) {
		if err = os.MkdirAll(dstDir, 0777); err != nil {
			return "", "", err
		}
	} else if err != nil {
		return "", "", err
	}

	// check for optional copy mode - default to all
	if len(os.Args) >= 4 {
		if err = validateMode(); err != nil {
			return "", "", err
		}
	}

	return srcDir, dstDir, nil
}

func validateMode() error {
	execMode = strings.ToLower(os.Args[3])

	// TODO: support custom mode (user supplied extension list)
	// check special modes
	switch execMode {
	case "all":
		return nil
	default:
		break
	}

	// check extension modes
	if _, ok := supportedExtensions[execMode]; !ok {
		return fmt.Errorf("mode not supported")
	}

	return nil
}
