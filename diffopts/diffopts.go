package diffopts

import (
	"fmt"
)

type Option struct {
	name           string // TODO rm it later
	FilterJSON     func([]byte) ([]byte, error)
	FilterLineDiff func(string) string
}

func IgnorePaths([]string) Option {
	return Option{
		name: "IgnroePath",
		FilterJSON: func(b []byte) ([]byte, error) {
			// TODO
			return b, nil
		},
		FilterLineDiff: func(line string) string { return line },
	}
}

func Colorize() Option {
	return Option{
		name:       "Colorize",
		FilterJSON: func(b []byte) ([]byte, error) { return b, nil },
		FilterLineDiff: func(line string) string {
			switch line[:1] {
			case "+":
				line = fmt.Sprintf("\x1b[32m%s\x1b[0m", line)
			case "-":
				line = fmt.Sprintf("\x1b[31m%s\x1b[0m", line)
			}
			return line
		},
	}
}
