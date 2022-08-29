package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github/venkat/twitter/x/twitter/types"
)

func (k msgServer) FetchFeed(goCtx context.Context, msg *types.MsgFetchFeed) (*types.MsgFetchFeedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userData, found := k.Keeper.GetUserData(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCantFetchUserData, "User %s not found", msg.Creator)
	}

	var followerTweets []string
	for _, follower := range userData.Followers {
		followerData, found := k.Keeper.GetUserData(ctx, follower)
		if !found {
			continue
		}
		for _, tweet := range followerData.Tweets {
			followerTweets = append(followerTweets, tweet)
		}
	}

	userFeed, found := k.Keeper.GetUserFeed(ctx, msg.Creator)
	if !found {
		userFeed.Index = msg.Creator
	}
	userFeed.Feed = followerTweets
	k.Keeper.SetUserFeed(ctx, userFeed)

	return &types.MsgFetchFeedResponse{
		FollowerTweets: followerTweets,
	}, nil
}
