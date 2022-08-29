package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddFollower = "add_follower"

var _ sdk.Msg = &MsgAddFollower{}

func NewMsgAddFollower(creator string, followerId string) *MsgAddFollower {
	return &MsgAddFollower{
		Creator:    creator,
		FollowerId: followerId,
	}
}

func (msg *MsgAddFollower) Route() string {
	return RouterKey
}

func (msg *MsgAddFollower) Type() string {
	return TypeMsgAddFollower
}

func (msg *MsgAddFollower) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddFollower) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddFollower) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
