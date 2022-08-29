package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddTweet = "add_tweet"

var _ sdk.Msg = &MsgAddTweet{}

func NewMsgAddTweet(creator string, content string) *MsgAddTweet {
	return &MsgAddTweet{
		Creator: creator,
		Content: content,
	}
}

func (msg *MsgAddTweet) Route() string {
	return RouterKey
}

func (msg *MsgAddTweet) Type() string {
	return TypeMsgAddTweet
}

func (msg *MsgAddTweet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddTweet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddTweet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
