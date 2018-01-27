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

func TestEqual(t *testing.T) {
	assert := assert.New(t)
	js := `{
		"foo": [1, 2],
		"bar": [3]
	}`
	for _, tt := range []struct {
		name string
		args string
		want bool
	}{
		{
			name: "equal",
			args: `{
				"bar": [3],
				"foo": [1, 2]
			}`,
			want: true,
		},
		{
			name: "not equal",
			args: `{
				"foo": [2, 1],
				"bar": [3]
			}`,
			want: false,
		},
		{
			name: "invalid json",
			args: `{
				"foo": [1, 2],
				"bar": [3],
			}`,
			want: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(
				tt.want, Equal([]byte(js), []byte(tt.args)),
			)
		})
	}
}

func TestDiff(t *testing.T) {
	assert := assert.New(t)
	js := `{
		"foo": [1, 2],
		"bar": [3]
	}`
	for _, tt := range []struct {
		name string
		args string
		want string
	}{
		{
			name: "equal",
			args: `{
				"bar": [3],
				"foo": [1, 2]
			}`,
			want: "",
		},
		{
			name: "not equal",
			args: `{
				"foo": [2, 1],
				"bar": [3]
			}`,
			want: `  {
    "bar": [
      3
    ],
    "foo": [
-     1,
-     2
+     2,
+     1
    ]
  }`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(
				tt.want, Diff([]byte(js), []byte(tt.args)),
			)
		})
	}
}
