package loadstream

import (
	"flag"
	"path"
	"testing"
)

var testFilesFolderPath = flag.String("path", "", "path to mock file")

func BenchmarkLoadStreamFile_sample2K(b *testing.B) {
	var sample2K = path.Join(*testFilesFolderPath, "sample2K.log")
	for i := 0; i < b.N; i++ {
		LoadStreamFile(sample2K, func(a ...interface{}) (n int, err error) { return })
	}
}

func BenchmarkLoadStreamFile_sample20K(b *testing.B) {
	var sample2K = path.Join(*testFilesFolderPath, "sample20K.log")
	for i := 0; i < b.N; i++ {
		LoadStreamFile(sample2K, func(a ...interface{}) (n int, err error) { return })
	}
}

func BenchmarkLoadStreamFile_sample200K(b *testing.B) {
	var sample2K = path.Join(*testFilesFolderPath, "sample200K.log")
	for i := 0; i < b.N; i++ {
		LoadStreamFile(sample2K, func(a ...interface{}) (n int, err error) { return })
	}
}
