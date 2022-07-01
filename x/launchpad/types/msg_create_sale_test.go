package types

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestValidateMsgCreateSale(t *testing.T) {
	minDuration := 2 * time.Hour
	minUntilSale := time.Hour
	now := time.Now()
	_, _, creatorAddr := testdata.KeyTestPubAddr()
	creator := creatorAddr.String()
	// creatorBad := sdk.AccAddress("invalid").String()
	tOut := sdk.NewCoin("tokenout", sdk.NewIntFromUint64(1000))
	url1 := "https://api.ipfsbrowser.com/ipfs/get.php?hash=QmcGV8fimB7aeBxnDqr7bSSLUWLeyFKUukGqDhWnvriQ3T"
	tcs := []struct {
		name   string
		errMsg string
		m      MsgCreateSale
	}{
		// TODO: add more tests
		{"good1", "",
			MsgCreateSale{Creator: creator, TokenIn: "osmo", TokenOut: &tOut, StartTime: now.Add(minUntilSale), Duration: minDuration, Name: "My token sale", Url: url1}},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			require := require.New(t)
			c, err := tc.m.Validate(now, minDuration, minUntilSale)
			if tc.errMsg == "" {
				require.NoError(err)
				require.Equal(tc.m.Creator, c.String())
			} else {
				require.ErrorContains(err, tc.errMsg)
			}
		})
	}
}
