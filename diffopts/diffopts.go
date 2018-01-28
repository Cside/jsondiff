package diffopts

import (
	"fmt"

	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch"
)

type Option struct {
	name           string // TODO rm it later
	FilterJSON     func([]byte) ([]byte, error)
	FilterLineDiff func(string) string
}

type Patch struct {
	Op   string `json:"op"`
	Path string `json:"path"`
}

func IgnorePaths(paths []string) Option {
	p := []Patch{}
	for _, path := range paths {
		p = append(p, Patch{Op: "remove", Path: path})
	}
	pb, _ := json.Marshal(&p)
	patch, _ := jsonpatch.DecodePatch(pb)

	return Option{
		name: "IgnroePath",
		FilterJSON: func(b []byte) ([]byte, error) {
			return patch.Apply(b)
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
