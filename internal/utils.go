package internal

import "fmt"

// TODO: pass in the different sets of extensions supported
func getRegex() string {
	r := matchRegex
	first := true

	for _, ext := range imageExtensions {
		if first {
			first = false
		} else {
			r += "|"
		}

		r += fmt.Sprintf("(.*%s$)", ext)
	}

	return r
}
