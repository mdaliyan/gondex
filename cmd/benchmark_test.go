package main

import (
	"os"
	"testing"

	"github.com/mdaliyan/gondex/cmd/internal/crawler"
)

func BenchmarkOsDir(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ioutilReadDir()
	}
}

var files []os.FileInfo

func ioutilReadDir() {
	crawler.Generate(tmp)
}
