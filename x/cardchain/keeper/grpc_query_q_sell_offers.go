package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QSellOffers(goCtx context.Context, req *types.QueryQSellOffersRequest) (*types.QueryQSellOffersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var (
		sellOfferIds   []uint64
		sellOffersList []*types.SellOffer
	)

	iter := k.SellOffers.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, sellOffer := iter.Value()
		// Checks for price
		if !req.Ignore.Price {
			// Conversion to coins
			priceUp, err := sdk.ParseCoinNormalized(req.PriceUp)
			if err != nil {
				return nil, sdkerrors.Wrap(types.ErrConversion, err.Error())
			}
			priceDown, err := sdk.ParseCoinNormalized(req.PriceDown)
			if err != nil {
				return nil, sdkerrors.Wrap(types.ErrConversion, err.Error())
			}

			if !(sellOffer.Price.IsGTE(priceDown) && priceUp.IsGTE(sellOffer.Price)) {
				continue
			}
		}

		// Checks for seller
		if !req.Ignore.Seller {
			if req.Seller != sellOffer.Seller {
				continue
			}
		}

		// Checks for buyer
		if !req.Ignore.Buyer {
			if req.Seller != sellOffer.Buyer {
				continue
			}
		}

		// Checks for Status
		if !req.Ignore.Status {
			if req.Status != sellOffer.Status {
				continue
			}
		}

		// Checks for Card
		if !req.Ignore.Card {
			if req.Card != sellOffer.Card {
				continue
			}
		}

		sellOffersList = append(sellOffersList, sellOffer)
		sellOfferIds = append(sellOfferIds, uint64(idx))
	}

	return &types.QueryQSellOffersResponse{sellOfferIds, sellOffersList}, nil
}
