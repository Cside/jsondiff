# jsondiff [![wercker status](https://app.wercker.com/status/cf6fb86dc49eadcac25f466f8df9d05b/s/master "wercker status")](https://app.wercker.com/project/byKey/cf6fb86dc49eadcac25f466f8df9d05b) [![GoDoc](https://godoc.org/github.com/Cside/jsondiff?status.svg)](http://godoc.org/github.com/Cside/jsondiff)

Reports differences between two JSONs

## Example

```go
import (
	"testing"
	"github.com/Cside/jsondiff"
)

a := []byte(`{ "foo":1, "bar":2 }`)
b := []byte(`{ "foo":1, "bar":3 }`)

if diff := jsondiff.Diff(a, b); diff != "" {
	t.Errorf("two jsons are not equal. diff:\n%s", diff)
}
```

output:

```
=== RUN   TestMain
--- FAIL: TestMain (0.00s)
        main_test.go:14: two jsons are not equal. diff:
                  {
                -   "bar": 2,
                +   "bar": 3,
                    "foo": 1
                  }
FAIL
exit status 1
```

## Ignore values

You can ignore some values with specifiying `diffopts.IgnorePath(paths []string)` option.

Target values can be specified with [JSON Pointer](https://tools.ietf.org/html/rfc6901) .

The following test passes.

```go
import (
	"testing"
	"github.com/Cside/jsondiff"
	"github.com/Cside/jsondiff/diffopts"
)

a := []byte(`{ "foo": 1, "createdAt": 1517141881 }`)
b := []byte(`{ "foo": 1, "createdAt": 1528845681 }`)

if diff := jsondiff.Diff(a, b, diffopts.IgnorePaths([]string{"/createdAt"})); diff != "" {
	t.Errorf("two jsons are not equal. diff:\n%s", diff)
}
```

## License

[The MIT License](/LICENSE).
