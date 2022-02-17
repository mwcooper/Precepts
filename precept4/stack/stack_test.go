// Benchmark  cases for fib.go
package stack

import (
	"testing"
)

func BenchmarkStack20(b *testing.B) {
	ItemStack stack = 
	for n := 0; n < b.N; n++ {
		for items := 0; items < 20; items++ {
			stack(20) // push then pop 20 items b.N times
		}
	}
}
