package mustNot

import (
	"os"
	"strings"

	"github.com/klauspost/cpuid"
)

func init() {
	if !cpuid.CPU.Intel() {
		return
	}

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-test.bench=") {
			panic("you are on a fucking Intel processor " +
				"and must not run benchmarks")
		}
	}
}
