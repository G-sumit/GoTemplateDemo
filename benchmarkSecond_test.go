package main

import (
	"testing"
)

func BenchmarkDirectConversion(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conversionDirect()
	}
}