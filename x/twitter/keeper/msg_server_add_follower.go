package keeper

import (
	"context"

	"github/venkat/twitter/x/twitter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddFollower(goCtx context.Context, msg *types.MsgAddFollower) (*types.MsgAddFollowerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userData, found := k.Keeper.GetUserData(ctx, msg.Creator)
	if !found {
		userData.Index = msg.Creator
	}
	userData.Followers = append(userData.Followers, msg.FollowerId)
	k.Keeper.SetUserData(ctx, userData)

	return &types.MsgAddFollowerResponse{}, nil
}
