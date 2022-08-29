package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github/venkat/twitter/x/twitter/types"
)

// SetUserFeed set a specific userFeed in the store from its index
func (k Keeper) SetUserFeed(ctx sdk.Context, userFeed types.UserFeed) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserFeedKeyPrefix))
	b := k.cdc.MustMarshal(&userFeed)
	store.Set(types.UserFeedKey(
		userFeed.Index,
	), b)
}

// GetUserFeed returns a userFeed from its index
func (k Keeper) GetUserFeed(
	ctx sdk.Context,
	index string,

) (val types.UserFeed, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserFeedKeyPrefix))

	b := store.Get(types.UserFeedKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserFeed removes a userFeed from the store
func (k Keeper) RemoveUserFeed(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserFeedKeyPrefix))
	store.Delete(types.UserFeedKey(
		index,
	))
}

// GetAllUserFeed returns all userFeed
func (k Keeper) GetAllUserFeed(ctx sdk.Context) (list []types.UserFeed) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserFeedKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserFeed
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
