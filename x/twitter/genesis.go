package twitter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github/venkat/twitter/x/twitter/keeper"
	"github/venkat/twitter/x/twitter/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the userData
	for _, elem := range genState.UserDataList {
		k.SetUserData(ctx, elem)
	}
	// Set all the userFeed
	for _, elem := range genState.UserFeedList {
		k.SetUserFeed(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.UserDataList = k.GetAllUserData(ctx)
	genesis.UserFeedList = k.GetAllUserFeed(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
