package transfer

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

var (
	supportedExtensions = map[string][]string{
		"image":        {"jpg", "jpeg", "png", "aae", "heic", "raw", "mpo", "nef"},
		"video":        {"mp4", "mov", "flv", "3gp", "mpg", "wmv", "avi", "webm"},
		"audio":        {"mp3", "wav", "ogg", "m4a", "flac", "wma", "aac"},
		"pdf":          {"pdf"},
		"text":         {"txt", "docx", "doc", "rtf"},
		"presentation": {"pptx", "ppt"},
		"spreadsheet":  {"xls", "xlsx"},
		"archive":      {"zip", "rar", "7zip"},
	}
)

// runtime vars
var (
	rgx *regexp.Regexp
	pb  *progressbar.ProgressBar

	execMode = "all"
)

// job vars
var (
	totalFiles int
	fileMu     sync.Mutex
)
