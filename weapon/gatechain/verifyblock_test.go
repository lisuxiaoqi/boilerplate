package gatechain

import (
	"encoding/hex"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBlockVerification(t *testing.T) {
	rawBytes, err := hex.DecodeString(certStr)
	require.NoError(t, err)

	bcert, err := decodeBlockCert(t, rawBytes)
	require.NoError(t, err)

	t.Log("block decoded",
		"\nheight", bcert.Block.BlockHeader.Round,
		"\nblock hash pretty\t", bcert.Block.Hash().String(),
		"\nblock hash\t", bcert.Block.Digest().String(),
		"\ncert hash\t", bcert.Certificate.Proposal.BlockDigest.String(),
		"\ncert len\t", len(bcert.Certificate.Votes))
}
