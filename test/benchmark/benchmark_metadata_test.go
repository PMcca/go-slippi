package benchmark_test

import (
	goslippi "github.com/pmcca/go-slippi"
	"github.com/pmcca/go-slippi/slippi"
	"testing"
)

// Ensure compiler doesn't optimise-out function call, by setting results to these pkg-level variables.
var result slippi.Metadata
var err error

func BenchmarkMetadataOnly(b *testing.B) {
	var r slippi.Metadata
	var e error
	filePath := "../replays/metadata.slp"
	for i := 0; i < b.N; i++ {
		r, err = goslippi.ParseMeta(filePath)
		if err != nil {
			b.Fatal(err)
		}
	}
	result = r
	err = e
}
