package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github/venkat/twitter/testutil/keeper"
	"github/venkat/twitter/x/twitter/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TwitterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
