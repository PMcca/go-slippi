package test

import (
	goslippi "github.com/pmcca/go-slippi"
	"github.com/pmcca/go-slippi/slippi"
	"github.com/stretchr/testify/require"
	"testing"
)

// mustParseSlippiGame parses and returns a Slippi game, or fails the test if an error occurred.
func mustParseSlippiGame(t *testing.T, filePath string) slippi.Game {
	actual, err := goslippi.ParseGame(filePath)
	require.NoError(t, err)
	return actual
}

// mustParseSlippiMeta parses and returns a Slippi Metadata, or fails the test if an error occurred.
func mustParseSlippiMeta(t *testing.T, filePath string) slippi.Metadata {
	actual, err := goslippi.ParseMeta(filePath)
	require.NoError(t, err)
	return actual
}
