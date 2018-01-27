package jsondiff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineDiff(t *testing.T) {
	assert := assert.New(t)
	str := "foo\n" +
		"bar\n" +
		"\n" +
		"baz"
	for _, tt := range []struct {
		name string
		args string
		want string
	}{
		{
			name: "equal",
			args: "foo\n" +
				"bar\n" +
				"\n" +
				"baz",
			want: "",
		},
		{
			name: "plus",
			args: "foo\n" +
				"bar\n" +
				"\n" +
				"foobar\n" +
				"baz",
			want: "  foo\n" +
				"  bar\n" +
				"  \n" +
				"+ foobar\n" +
				"  baz",
		},
		{
			name: "minus",
			args: "foo\n" +
				"bar\n" +
				"baz",
			want: "  foo\n" +
				"  bar\n" +
				"- \n" +
				"  baz",
		},
		{
			name: "plus and minus",
			args: "foo\n" +
				"bar\n" +
				"foobar\n" +
				"baz",
			want: "  foo\n" +
				"  bar\n" +
				"- \n" +
				"+ foobar\n" +
				"  baz",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.want, LineDiff(str, tt.args))
		})
	}
}
