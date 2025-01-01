package internal

import "regexp"

type RecursiveImageCopyJob struct {
	ID  int
	Dir string

	regex *regexp.Regexp
}

// implement
func (j *RecursiveImageCopyJob) Process() {

}
