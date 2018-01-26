package jsondiff

import (
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func Equal(a, b []byte) bool {
	return true
}

func Diff(a, b []byte) string {
	return ""
}

var linePrefix = map[diffmatchpatch.Operation]string{
	diffmatchpatch.DiffEqual:  "  ",
	diffmatchpatch.DiffInsert: "+ ",
	diffmatchpatch.DiffDelete: "- ",
}

// TODO: color option
func LineDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	a, b, c := dmp.DiffLinesToChars(a, b)
	diffs := dmp.DiffMain(a, b, false)

	diffStrs := []string{}
	modified := 0
	for _, diff := range dmp.DiffCharsToLines(diffs, c) {
		if diff.Type != diffmatchpatch.DiffEqual {
			modified++
		}
		texts := strings.Split(diff.Text, "\n")
		for _, text := range texts {
			if text == "" {
				continue
			}
			diffStrs = append(diffStrs, linePrefix[diff.Type]+text)
		}
	}
	if modified == 0 {
		return ""
	}
	return strings.Join(diffStrs, "\n")
}
