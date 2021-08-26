package main

import (
	"testing"
)

func BenchmarkGoTemplateConversion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoTemplate()
	}
}