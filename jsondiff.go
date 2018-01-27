package jsondiff

import (
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/fatih/color"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var linePrefix = map[diffmatchpatch.Operation]string{
	diffmatchpatch.DiffEqual:  "  ",
	diffmatchpatch.DiffInsert: "+ ",
	diffmatchpatch.DiffDelete: "- ",
}

type Option struct {
	filterJSON     func([]byte) ([]byte, error)
	filterLineDiff func(string) string
}

func IgnorePaths([]string) Option {
	return Option{
		filterJSON: func(b []byte) ([]byte, error) {
			// TODO
			return b, nil
		},
		filterLineDiff: func(line string) string { return line },
	}
}

func Colorize() Option {
	return Option{
		filterJSON: func(b []byte) ([]byte, error) { return b, nil },
		filterLineDiff: func(line string) string {
			switch line[:1] {
			case "+":
				line = color.HiGreenString(line)
			case "-":
				line = color.RedString(line)
			}
			return line
		},
	}
}

func Equal(a, b []byte, opts ...Option) bool {
	return Diff(a, b, opts...) == ""
}

func Diff(a, b []byte, opts ...Option) string {
	return LineDiff(
		string(BeautifyJSON(a, opts...)),
		string(BeautifyJSON(b, opts...)),
		opts...,
	)
}

// TODO ignore path
func BeautifyJSON(b []byte, opts ...Option) []byte {
	for _, opt := range opts {
		filtered, err := opt.filterJSON(b)
		if err == nil {
			b = filtered
		}
	}
	js, err := simplejson.NewJson(b)
	if err != nil {
		return []byte("invalid JSON")
	}
	out, err := js.EncodePretty()
	if err != nil {
		return []byte("invalid JSON")
	}
	return out
}

// TODO: color option
func LineDiff(a, b string, opts ...Option) string {
	dmp := diffmatchpatch.New()
	a, b, c := dmp.DiffLinesToChars(a, b)
	diffs := dmp.DiffMain(a, b, false)

	lines := []string{}
	modified := 0
	for _, diff := range dmp.DiffCharsToLines(diffs, c) {
		if diff.Type != diffmatchpatch.DiffEqual {
			modified++
		}
		texts := strings.Split(
			strings.TrimSuffix(diff.Text, "\n"),
			"\n",
		)
		for _, text := range texts {
			line := linePrefix[diff.Type] + text
			lines = append(lines, line)
		}
	}
	if modified == 0 {
		return ""
	}
	newLines := []string{}
	for _, line := range lines {
		for _, opt := range opts {
			line = opt.filterLineDiff(line)
		}
		newLines = append(newLines, line)
	}
	return strings.Join(newLines, "\n")
}
