package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github/venkat/twitter/testutil/keeper"
	"github/venkat/twitter/testutil/nullify"
	"github/venkat/twitter/x/twitter/keeper"
	"github/venkat/twitter/x/twitter/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUserFeed(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UserFeed {
	items := make([]types.UserFeed, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetUserFeed(ctx, items[i])
	}
	return items
}

func TestUserFeedGet(t *testing.T) {
	keeper, ctx := keepertest.TwitterKeeper(t)
	items := createNUserFeed(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUserFeed(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUserFeedRemove(t *testing.T) {
	keeper, ctx := keepertest.TwitterKeeper(t)
	items := createNUserFeed(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserFeed(ctx,
			item.Index,
		)
		_, found := keeper.GetUserFeed(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestUserFeedGetAll(t *testing.T) {
	keeper, ctx := keepertest.TwitterKeeper(t)
	items := createNUserFeed(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUserFeed(ctx)),
	)
}
