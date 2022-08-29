package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github/venkat/twitter/x/twitter/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserFeedAll(c context.Context, req *types.QueryAllUserFeedRequest) (*types.QueryAllUserFeedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userFeeds []types.UserFeed
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	userFeedStore := prefix.NewStore(store, types.KeyPrefix(types.UserFeedKeyPrefix))

	pageRes, err := query.Paginate(userFeedStore, req.Pagination, func(key []byte, value []byte) error {
		var userFeed types.UserFeed
		if err := k.cdc.Unmarshal(value, &userFeed); err != nil {
			return err
		}

		userFeeds = append(userFeeds, userFeed)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserFeedResponse{UserFeed: userFeeds, Pagination: pageRes}, nil
}

func (k Keeper) UserFeed(c context.Context, req *types.QueryGetUserFeedRequest) (*types.QueryGetUserFeedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetUserFeed(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUserFeedResponse{UserFeed: val}, nil
}
