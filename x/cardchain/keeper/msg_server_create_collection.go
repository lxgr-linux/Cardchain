package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateCollection(goCtx context.Context, msg *types.MsgCreateCollection) (*types.MsgCreateCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	collectionId := k.GetCollectionsNumber(ctx)

	err := k.CollectCollectionCreationFee(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	collection := types.Collection{
		Name:         msg.Name,
		Cards:        []uint64{},
		Contributors: append([]string{msg.Creator}, msg.Contributors...),
		Artist:       msg.Artist,
		StoryWriter:  msg.StoryWriter,
		Status:       types.CStatus_design,
		TimeStamp:    0,
	}

	k.SetCollection(ctx, collectionId, collection)

	return &types.MsgCreateCollectionResponse{}, nil
}
