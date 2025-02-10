package transfer

import "fmt"

func buildExtensionRegex() (string, error) {
	r := matchRegex
	first := true

	l, err := buildExtensionList()
	if err != nil {
		return "", err
	}

	for _, ext := range l {
		if first {
			first = false
		} else {
			r += "|"
		}

		r += fmt.Sprintf("(.*%s$)", ext)
	}

	return r, nil
}

func buildExtensionList() ([]string, error) {
	fmt.Printf("MODE: %s", execMode)
	if v, ok := supportedExtensions[execMode]; ok {
		return v, nil
	}

	// TODO: support custom mode
	switch execMode {
	default:
		return nil, fmt.Errorf("mode %s not supported", execMode)
	}
}
