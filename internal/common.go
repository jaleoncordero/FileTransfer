package internal

import (
	"regexp"
	"sync"

	"github.com/schollz/progressbar/v3"
)

const (
	matchRegex     = `(?i)`
	testDir        = "./internal/test"
	workerPoolSize = 100
)

// supported extensions
var (
	imageExtensions = []string{"jpg", "jpeg", "png", "aae", "heic", "raw", "mpo", "nef"}
	// videoExtensions = []string{"mp4", "mov"}
)

// runtime vars
var (
	rgx    *regexp.Regexp
	srcDir string
	dstDir string

	bar *progressbar.ProgressBar
)

// job vars
var (
	totalFiles int
	fileMu     sync.Mutex
)
