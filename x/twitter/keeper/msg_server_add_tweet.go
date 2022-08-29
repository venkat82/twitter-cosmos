package keeper

import (
	"context"

	"github/venkat/twitter/x/twitter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddTweet(goCtx context.Context, msg *types.MsgAddTweet) (*types.MsgAddTweetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userData, found := k.Keeper.GetUserData(ctx, msg.Creator)
	if !found {
		userData.Index = msg.Creator
	}
	userData.Tweets = append(userData.Tweets, msg.Content)
	k.Keeper.SetUserData(ctx, userData)

	return &types.MsgAddTweetResponse{}, nil
}
