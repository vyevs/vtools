package vtools

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ReadLines reads file fPath and returns all the non-empty lines.
func ReadLines(fPath string) ([]string, error) {
	f, err := os.Open(fPath)
	if err != nil {
		return nil, fmt.Errorf("os.Open failed: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	lines := make([]string, 0, 128)
	for sc.Scan() {
		line := sc.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("error reading lines: %v", err)
	}

	return lines, nil
}

// ReadLinesBytes returns all the lines of file fPath as bytes.
func ReadLinesBytes(fPath string) ([][]byte, error) {
	f, err := os.Open(fPath)
	if err != nil {
		return nil, fmt.Errorf("os.Open failed: %v", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	lines := make([][]byte, 0, 128)
	for {
		lineBytes, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				if len(lineBytes) > 1 {
					lines = append(lines, lineBytes)
				}
				break
			}

			return nil, fmt.Errorf("failed to read line: %v", err)
		}

		if len(lineBytes) > 1 {
			// bufio.Reader.ReadBytes returns the bytes up to and including the delimeter, so cut the newline off.
			lineBytes = lineBytes[:len(lineBytes)-1]
			lines = append(lines, lineBytes)
		}

	}

	return lines, nil
}
