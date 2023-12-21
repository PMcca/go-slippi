package test

import (
	goslippi "github.com/PMcca/go-slippi"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGeckoCodes(t *testing.T) {
	t.Parallel()
	filePath := "replays/geckoCodes.slp"
	actual, err := goslippi.ParseGame(filePath)
	require.NoError(t, err)

	expectedAddress := uint32(0x8015ee98)
	geckoList := actual.Data.GeckoCodes
	require.Len(t, geckoList.Codes, 457)
	require.Equal(t, expectedAddress, geckoList.Codes[0].Address)
}
