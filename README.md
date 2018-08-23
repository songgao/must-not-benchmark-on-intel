Intel
[says](https://perens.com/2018/08/22/new-intel-microcode-license-restriction-is-not-acceptable/)
you must not benchmark on their platform:

> You will not, and will not allow any third party to
> ...
> (v) publish or provide any Software benchmark or comparison test results.

This package helps make sure you never run any benchmarks on an Intel
processor.

# Usage

## Go

Just import `github.com/songgao/must-not-benchmark-on-intel/go` for
side-effects:

```go
import _ "github.com/songgao/must-not-benchmark-on-intel/go"
```

Example:

```
bash-3.2$ cat yo_test.go
package yo

import (
	"testing"

	_ "github.com/songgao/must-not-benchmark-on-intel/go"
)

func TestYo(*testing.T)      {}
func BenchmarkYo(*testing.B) {}
bash-3.2$ go test
PASS
ok  	_/tmp/t	0.009s
bash-3.2$ go test -bench='.*'
panic: Intel are sketchy fucks, therefor you must not run benchmarks on their CPUs.

goroutine 1 [running]:
github.com/songgao/must-not-benchmark-on-intel/go.init.0()
	/Users/songgao/gopath/src/github.com/songgao/must-not-benchmark-on-intel/go/must_not.go:17 +0x134
exit status 2
FAIL	_/tmp/t	0.010s
```

# License

[WTFPL](https://en.wikipedia.org/wiki/WTFPL)

# Contributing

Yes.
