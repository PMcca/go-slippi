package benchmark_test

import (
	goslippi "github.com/PMcca/go-slippi"
	"github.com/PMcca/go-slippi/slippi"
	"testing"
)

// Ensure compiler doesn't optimise-out function call, by setting results to these pkg-level variables.
var gameResult slippi.Game
var gameErr error

func BenchmarkGame(b *testing.B) {
	var r slippi.Game
	var e error
	filePath := "../replays/test2.slp"
	for i := 0; i < b.N; i++ {
		r, err = goslippi.ParseGame(filePath)
		if err != nil {
			b.Fatal(err)
		}
	}
	gameResult = r
	gameErr = e
}
