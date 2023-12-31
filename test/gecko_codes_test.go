package test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGeckoCodes(t *testing.T) {
	t.Parallel()

	actual := mustParseSlippiGame(t, "replays/geckoCodes.slp")
	expectedAddress := uint32(0x8015ee98)
	geckoList := actual.Data.GeckoCodes
	require.Len(t, geckoList.Codes, 457)
	require.Equal(t, expectedAddress, geckoList.Codes[0].Address)
}
