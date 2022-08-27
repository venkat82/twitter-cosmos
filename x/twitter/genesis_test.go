package twitter_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github/venkat/twitter/testutil/keeper"
	"github/venkat/twitter/testutil/nullify"
	"github/venkat/twitter/x/twitter"
	"github/venkat/twitter/x/twitter/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TwitterKeeper(t)
	twitter.InitGenesis(ctx, *k, genesisState)
	got := twitter.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
