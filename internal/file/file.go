package file

import (
	"bufio"
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
