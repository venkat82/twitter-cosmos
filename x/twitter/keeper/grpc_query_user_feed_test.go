package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github/venkat/twitter/testutil/keeper"
	"github/venkat/twitter/testutil/nullify"
	"github/venkat/twitter/x/twitter/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestUserFeedQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.TwitterKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNUserFeed(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetUserFeedRequest
		response *types.QueryGetUserFeedResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetUserFeedRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetUserFeedResponse{UserFeed: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetUserFeedRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetUserFeedResponse{UserFeed: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetUserFeedRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.UserFeed(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestUserFeedQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.TwitterKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNUserFeed(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllUserFeedRequest {
		return &types.QueryAllUserFeedRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UserFeedAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UserFeed), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UserFeed),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UserFeedAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UserFeed), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UserFeed),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.UserFeedAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.UserFeed),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.UserFeedAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
