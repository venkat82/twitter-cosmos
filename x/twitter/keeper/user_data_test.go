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

func createNUserData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UserData {
	items := make([]types.UserData, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetUserData(ctx, items[i])
	}
	return items
}

func TestUserDataGet(t *testing.T) {
	keeper, ctx := keepertest.TwitterKeeper(t)
	items := createNUserData(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUserData(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUserDataRemove(t *testing.T) {
	keeper, ctx := keepertest.TwitterKeeper(t)
	items := createNUserData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserData(ctx,
			item.Index,
		)
		_, found := keeper.GetUserData(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestUserDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.TwitterKeeper(t)
	items := createNUserData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUserData(ctx)),
	)
}
