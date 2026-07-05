package file

import (
	"bufio"
	"bytes"
	"iter"
	"os"
)

func ReadLines(name string) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		file, err := os.Open(name)
		if err != nil {
			yield("", err)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if !yield(scanner.Text(), nil) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			yield("", err)
		}
	}
}

func scanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := range data {
		if data[i] == ',' {
			return i + 1, data[:i], nil
		}
	}
	if atEOF {
		return len(data), bytes.TrimSpace(data), nil
	}
	return 0, nil, nil
}

func ReadCommas(fileName string) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		file, err := os.Open(fileName)
		if err != nil {
			yield("", err)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(scanCommas)
		for scanner.Scan() {
			if !yield(scanner.Text(), nil) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			yield("", err)
		}
	}
}
