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
			panic("Intel are sketchy fucks, therefor you " +
				"must not run benchmarks on their CPUs.")
		}
	}
}
