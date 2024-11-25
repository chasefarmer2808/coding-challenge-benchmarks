package input

import (
	"bufio"
	"iter"
	"os"
)

func Lines(f *os.File) iter.Seq[string] {
	return func(yield func(string) bool) {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}
}
