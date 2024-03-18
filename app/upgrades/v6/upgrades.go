package v6

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/sge-network/sge/app/keepers"
	markettypes "github.com/sge-network/sge/x/market/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	k *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		participationList, err := k.OrderbookKeeper.GetAllOrderBookParticipations(ctx)
		if err != nil {
			panic(err)
		}

		for _, bp := range participationList {
			if !bp.IsSettled {
				continue
			}

			market, found := k.MarketKeeper.GetMarket(ctx, bp.OrderBookUID)
			if !found {
				panic(fmt.Errorf("market not found %s", bp.OrderBookUID))
			}

			reimburseFee := false
			switch market.Status {
			case markettypes.MarketStatus_MARKET_STATUS_RESULT_DECLARED:
				bp.ReturnedAmount = bp.Liquidity.Add(bp.ActualProfit)
				if bp.NotParticipatedInBetFulfillment() {
					reimburseFee = true
				}

			case markettypes.MarketStatus_MARKET_STATUS_CANCELED,
				markettypes.MarketStatus_MARKET_STATUS_ABORTED:
				bp.ReturnedAmount = bp.Liquidity
				reimburseFee = true
			}

			if reimburseFee {
				bp.ReimbursedFee = bp.Fee
				bp.ReturnedAmount = bp.ReturnedAmount.Add(bp.ReimbursedFee)
			}

			k.OrderbookKeeper.SetOrderBookParticipation(ctx, bp)
		}

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
