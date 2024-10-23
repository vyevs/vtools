package vtools

import (
	"bufio"
	"fmt"
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

	lines := make([]string, 0, 32)
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
