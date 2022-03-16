package keeper

import (
	"context"

	"golang.org/x/exp/slices"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RemoveCardFromCollection(goCtx context.Context, msg *types.MsgRemoveCardFromCollection) (*types.MsgRemoveCardFromCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection := k.GetCollection(ctx, msg.CollectionId)
	if !slices.Contains(collection.Contributors, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid contributor")
	}
	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	newCards, err := PopItemFromArr(msg.CardId, collection.Cards)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "Card: %d", msg.CardId)
	}

	collection.Cards = newCards

	k.SetCollection(ctx, msg.CollectionId, collection)

	return &types.MsgRemoveCardFromCollectionResponse{}, nil
}
