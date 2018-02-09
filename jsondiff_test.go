package jsondiff

import (
	"testing"

	"fmt"

	"github.com/Cside/jsondiff/diffopts"
	"github.com/stretchr/testify/assert"
)

func b(s string) []byte { return []byte(s) }

func TestLineDiff(t *testing.T) {
	assert := assert.New(t)
	str := `foo
bar

baz`
	for _, tt := range []struct {
		name string
		args string
		want string
	}{
		{
			name: "equal",
			args: `foo
bar

baz`,
			want: ``,
		},
		{
			name: "plus",
			args: `foo
bar

foobar
baz`,
			want: "  foo\n" +
				"  bar\n" +
				"  \n" +
				"+ foobar\n" +
				"  baz",
		},
		{
			name: "minus",
			args: `foo
bar
baz`,
			want: "  foo\n" +
				"  bar\n" +
				"- \n" +
				"  baz",
		},
		{
			name: "plus and minus",
			args: `foo
bar
foobar
baz`,
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
				tt.want, Equal(b(js), b(tt.args)),
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
		{
			name: "js safe int",
			args: `{
				"foo": [1, 2],
				"bar": [9007199254740990]
			}`,
			want: `  {
    "bar": [
-     3
+     9007199254740990
    ],
    "foo": [
      1,
      2
    ]
  }`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(
				tt.want, Diff(b(js), b(tt.args)),
			)
		})
	}
}

func Test_diffoptsColorize(t *testing.T) {
	fmt.Println(
		Diff(
			b(`{ "foo": [1, 2], "bar": [3] }`),
			b(`{ "foo": [2, 1], "bar": [3] }`),
			diffopts.Colorize(),
		),
	)
}

func Test_diffoptsIgnorePaths(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		a      string
		b      string
		ignore []string
	}
	for _, tt := range []struct {
		name string
		args args
		want bool
	}{
		{
			name: "basic case",
			args: args{
				a:      `{"foo":1, "bar":2}`,
				b:      `{"foo":1, "bar":3}`,
				ignore: []string{"/bar"},
			},
			want: true,
		},
		{
			name: "multiple",
			args: args{
				a:      `{"foo":1, "bar":2, "baz":3}`,
				b:      `{"foo":1, "bar":3, "baz":4}`,
				ignore: []string{"/bar", "/baz"},
			},
			want: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(
				tt.want,
				Equal(b(tt.args.a), b(tt.args.b), diffopts.IgnorePaths(tt.args.ignore)),
			)
		})
	}
}
