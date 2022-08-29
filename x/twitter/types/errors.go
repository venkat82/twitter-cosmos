package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/twitter module sentinel errors
var (
	ErrCantFetchUserData = sdkerrors.Register(ModuleName, 1103, "cannot fetch user data for %s")
)
