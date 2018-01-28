package diffopts

import (
	"encoding/json"
	"fmt"

	"github.com/evanphx/json-patch"
)

type Option struct {
	FilterJSON     func([]byte) ([]byte, error)
	FilterLineDiff func(string) string
}

type patch struct {
	Op   string `json:"op"`
	Path string `json:"path"`
}

func IgnorePaths(paths []string) Option {
	p := []patch{}
	for _, path := range paths {
		p = append(p, patch{Op: "remove", Path: path})
	}
	pb, _ := json.Marshal(&p)
	patch, _ := jsonpatch.DecodePatch(pb)

	return Option{
		FilterJSON: func(b []byte) ([]byte, error) {
			return patch.Apply(b)
		},
		FilterLineDiff: func(line string) string { return line },
	}
}

func Colorize() Option {
	return Option{
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
