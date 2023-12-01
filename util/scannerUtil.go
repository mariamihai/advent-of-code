package util

import "strings"

func CustomSplit(customDataFunc func(data string) []byte) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// Return nothing if at end of file and no data passed
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		// Find the index of the input of a newline
		if i := strings.Index(string(data), "\n"); i >= 0 {
			return i + 1, customDataFunc(string(data[0:i])), nil
		}

		// If at end of file with data return the data
		if atEOF {
			return len(data), customDataFunc(string(data)), nil
		}

		return
	}
}
