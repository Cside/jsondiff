package jsondiff

import (
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func Equal(a, b []byte) bool {
	return string(BeautifyJSON(a)) == string(BeautifyJSON(b))
}

func Diff(a, b []byte) string {
	return LineDiff(
		string(BeautifyJSON(a)),
		string(BeautifyJSON(b)),
	)
}

// TODO ignore path
func BeautifyJSON(b []byte) []byte {
	js, err := simplejson.NewJson(b)
	if err != nil {
		return []byte("invalid JSON")
	}
	o, err := js.EncodePretty()
	if err != nil {
		return []byte("invalid JSON")
	}
	return o
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
		texts := strings.Split(
			strings.TrimSuffix(diff.Text, "\n"),
			"\n",
		)
		for _, text := range texts {
			diffStrs = append(diffStrs, linePrefix[diff.Type]+text)
		}
	}
	if modified == 0 {
		return ""
	}
	return strings.Join(diffStrs, "\n")
}
