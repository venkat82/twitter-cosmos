package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFetchFeed = "fetch_feed"

var _ sdk.Msg = &MsgFetchFeed{}

func NewMsgFetchFeed(creator string) *MsgFetchFeed {
	return &MsgFetchFeed{
		Creator: creator,
	}
}

func (msg *MsgFetchFeed) Route() string {
	return RouterKey
}

func (msg *MsgFetchFeed) Type() string {
	return TypeMsgFetchFeed
}

func (msg *MsgFetchFeed) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFetchFeed) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFetchFeed) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
