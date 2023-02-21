package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	bettypes "github.com/sge-network/sge/x/bet/types"
	"github.com/sge-network/sge/x/orderbook/types"
)

// ProcessBetPlacement processes bet placement
func (k Keeper) ProcessBetPlacement(
	ctx sdk.Context,
	betUID, bookID, oddsID string,
	maxLossMultiplier sdk.Dec,
	payoutProfit sdk.Int,
	bettorAddress sdk.AccAddress,
	betFee sdk.Int,
	betAmount sdk.Int,
	oddsType bettypes.OddsType,
	oddsVal string, betID uint64,
) (betFulfillments []*bettypes.BetFulfillment, err error) {
	// betFulfillments := []*bettypes.BetFulfillment{}

	// If lock exists, return error
	// Lock already exists means the bet is already placed for the given bet-uid
	if k.payoutLockExists(ctx, betUID) {
		err = sdkerrors.Wrapf(types.ErrLockAlreadyExists, "%s", betUID)
		return
	}

	// get book data by its id
	book, found := k.GetBook(ctx, bookID)
	if !found {
		err = sdkerrors.Wrapf(types.ErrOrderBookNotFound, "%s", bookID)
		return
	}
	// get exposures of the odds user is placing the bet
	bookExposure, found := k.GetBookOddsExposure(ctx, bookID, oddsID)
	if !found {
		err = sdkerrors.Wrapf(types.ErrOrderBookExposureNotFound, "%s , %s", bookID, oddsID)
		return
	}

	// make maps for participation and its exposures
	participationMap := make(map[uint64]types.BookParticipation)
	participationExposureMap := make(map[uint64]types.ParticipationExposure)

	bps, err := k.GetParticipationsOfBook(ctx, book.ID)
	if err != nil {
		return
	}
	if int(book.ParticipationCount) != len(bps) {
		err = sdkerrors.Wrapf(types.ErrBookParticipationsNotFound, "%s", bookID)
		return
	}
	for _, bp := range bps {
		participationMap[bp.Index] = bp
	}

	pes, err := k.GetExposureByBookAndOdds(ctx, bookID, oddsID)
	if err != nil {
		return
	}
	if int(book.ParticipationCount) != len(pes) {
		err = sdkerrors.Wrapf(types.ErrParticipationExposuresNotFound, "%s, %s", bookID, oddsID)
		return
	}
	for _, pe := range pes {
		participationExposureMap[pe.ParticipationIndex] = pe
	}

	remainingPayoutProfit, updatedFulfillmentQueue := payoutProfit, bookExposure.FulfillmentQueue
	liquidatedBetAmount := sdk.ZeroInt()

	// continue until updatedFulfillmentQueue gets empty
	for len(updatedFulfillmentQueue) > 0 {
		var participation types.BookParticipation
		var participationExposure types.ParticipationExposure
		participationIndex := updatedFulfillmentQueue[0]

		getQueueItemParticipation := func() error {
			participation, found = participationMap[participationIndex]
			if !found {
				return sdkerrors.Wrapf(types.ErrBookParticipationNotFound, "%s, %d", bookID, participationIndex)
			}
			participationExposure, found = participationExposureMap[participationIndex]
			if !found {
				return sdkerrors.Wrapf(types.ErrParticipationExposureNotFound, "%s, %d", bookID, participationIndex)
			}
			if participationExposure.IsFulfilled {
				return sdkerrors.Wrapf(types.ErrParticipationExposureAlreadyFilled, "%s, %d", bookID, participationIndex)
			}
			return nil
		}

		removeQueueItem := func() {
			updatedFulfillmentQueue = updatedFulfillmentQueue[1:]
		}

		setFulfilled := func() {
			participationExposure.IsFulfilled = true
			participation.ExposuresNotFilled--
		}

		processBetFulfillment := func(amount sdk.Int, noMoreLiquidity bool) error {
			participationExposure.Exposure = participationExposure.Exposure.Add(amount)
			betAmount, err := bettypes.CalculateBetAmount(oddsType, oddsVal, amount)
			if err != nil {
				return err
			}

			liquidatedBetAmount = liquidatedBetAmount.Add(betAmount)
			participationExposure.BetAmount = participationExposure.BetAmount.Add(betAmount)
			if noMoreLiquidity {
				setFulfilled()
			}
			participation.TotalBetAmount = participation.TotalBetAmount.Add(betAmount)
			participation.CurrentRoundTotalBetAmount = participation.CurrentRoundTotalBetAmount.Add(betAmount)

			maxLoss := participationExposure.Exposure.Sub(participation.CurrentRoundTotalBetAmount).Add(participationExposure.BetAmount)
			switch {
			case participation.CurrentRoundMaxLoss.IsNil():
				participation.CurrentRoundMaxLoss = maxLoss
				participation.CurrentRoundMaxLossOddsID = oddsID
			case participation.CurrentRoundMaxLossOddsID == oddsID:
				participation.CurrentRoundMaxLoss = maxLoss
			default:
				originalMaxLoss := participation.CurrentRoundMaxLoss.Sub(betAmount)
				if maxLoss.GT(originalMaxLoss) {
					participation.CurrentRoundMaxLoss = maxLoss
					participation.CurrentRoundMaxLossOddsID = oddsID
				} else {
					participation.CurrentRoundMaxLoss = originalMaxLoss
				}
			}

			betFulfillments = append(betFulfillments, &bettypes.BetFulfillment{
				ParticipantAddress: participation.ParticipantAddress,
				ParticipationIndex: participation.Index,
				BetAmount:          betAmount,
				PayoutAmount:       amount,
			})
			remainingPayoutProfit = remainingPayoutProfit.Sub(amount)
			participationBetPair := types.NewParticipationBetPair(participation.BookID, betUID, participation.Index)
			k.SetParticipationBetPair(ctx, participationBetPair, betID)
			return nil
		}

		refreshQueueAndState := func() (err error) {
			maxLoss := sdk.MaxInt(sdk.ZeroInt(), participation.CurrentRoundMaxLoss)
			participation.CurrentRoundLiquidity = participation.CurrentRoundLiquidity.Sub(maxLoss)

			// check if there is more liquidity amount
			eligibleForNextRound := participation.CurrentRoundLiquidity.GT(sdk.ZeroInt())

			participationExposures, err := k.GetExposureByBookAndParticipationIndex(ctx, bookID, participationIndex)
			if err != nil {
				return
			}
			for _, pe := range participationExposures {
				k.MoveToHistoricalParticipationExposure(ctx, pe)
				if eligibleForNextRound {
					newPe := types.NewParticipationExposure(book.ID, pe.OddsID, sdk.ZeroInt(), sdk.ZeroInt(), pe.ParticipationIndex, pe.Round+1, false)
					k.SetParticipationExposure(ctx, newPe)
					if pe.OddsID == participationExposure.OddsID {
						participationExposure = newPe
						participationExposureMap[pe.ParticipationIndex] = participationExposure
					}
				}
			}

			participation.ExposuresNotFilled = book.OddsCount
			participation.CurrentRoundTotalBetAmount = sdk.ZeroInt()
			participation.MaxLoss = participation.MaxLoss.Add(participation.CurrentRoundMaxLoss)
			participation.CurrentRoundMaxLoss = sdk.ZeroInt()
			participationMap[participation.Index] = participation
			k.SetBookParticipation(ctx, participation)

			if eligibleForNextRound {
				var boes []types.BookOddsExposure
				boes, err = k.GetOddsExposuresByBook(ctx, bookID)
				if err != nil {
					return
				}
				for _, boe := range boes {
					boe.FulfillmentQueue = append(boe.FulfillmentQueue, participationIndex)
					if boe.OddsID == participationExposure.OddsID {
						bookExposure = boe
					}

					k.SetBookOddsExposure(ctx, boe)
				}
				updatedFulfillmentQueue = append(updatedFulfillmentQueue, participationIndex)
				bookExposure.FulfillmentQueue = updatedFulfillmentQueue
			}

			return nil
		}

		// fill participation and exposure values
		err = getQueueItemParticipation()
		if err != nil {
			return
		}

		// availableLiquidty is the available amount of tokens to be used from the participation exposure
		availableLiquidty := maxLossMultiplier.
			MulInt(participation.CurrentRoundLiquidity).TruncateInt().
			Sub(participationExposure.Exposure)

		switch {
		case availableLiquidty.LTE(sdk.ZeroInt()):
			setFulfilled()
			removeQueueItem()
		case availableLiquidty.LTE(remainingPayoutProfit):
			// if the available liquidity is less than remaining payout profit that
			// need to be paid, we should use all of available liquidity pull for the calculations.
			err := processBetFulfillment(availableLiquidty, true)
			if err != nil {
				return betFulfillments, err
			}
			removeQueueItem()
		default:
			// availableLiquidty is positive and more than remaining payout profit that
			// need to be paid, so we can cover all of payout profits with available liquidity.
			err := processBetFulfillment(remainingPayoutProfit, false)
			if err != nil {
				return betFulfillments, err
			}
		}

		k.SetParticipationExposure(ctx, participationExposure)
		k.SetBookParticipation(ctx, participation)

		// if there are no more exposures to be filled
		if participation.ExposuresNotFilled == 0 {
			refreshQueueAndState()
		}

		if remainingPayoutProfit.LTE(sdk.ZeroInt()) || len(updatedFulfillmentQueue) == 0 {
			break
		}
	}

	if remainingPayoutProfit.GT(sdk.ZeroInt()) {
		return betFulfillments, sdkerrors.Wrapf(types.ErrInternalProcessingBet, "insufficient liquidity in order book")
	}

	bookExposure.FulfillmentQueue = updatedFulfillmentQueue
	k.SetBookOddsExposure(ctx, bookExposure)

	// Transfer bet fee from bettor to the `bet` module account
	err = k.transferFundsFromUserToModule(ctx, bettorAddress, bettypes.ModuleName, betFee)
	if err != nil {
		return betFulfillments, err
	}

	// Transfer bet amount from bettor to `book_liquidity_pool` Account
	err = k.transferFundsFromUserToModule(ctx, bettorAddress, types.BookLiquidityName, liquidatedBetAmount)
	if err != nil {
		return betFulfillments, err
	}

	// Create a unique lock in the Payout Store for the bet
	k.setPayoutLock(ctx, betUID)

	return betFulfillments, nil
}